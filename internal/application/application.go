package application

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gateway_status "github.com/lallison/h_skills_project/internal/service/gateways/status"
	repository_status "github.com/lallison/h_skills_project/internal/service/repository/status"
	usecase_status "github.com/lallison/h_skills_project/internal/service/usecase/status"
	"github.com/lallison/h_skills_project/version"
)

type Application struct {
	Host, Port string
	srv        *http.Server

	ctx       context.Context
	ctxCancel context.CancelFunc

	log *slog.Logger
}

func New() *Application {
	mux := http.NewServeMux()

	// вынести в cmd
	// сделать метод RegisterHandlers(), которые принимает gateways
	repository := repository_status.New()
	useCase := usecase_status.New(repository)
	gateways := gateway_status.New(useCase)

	mux.HandleFunc("/status", gateways.HandleGetStatus)

	srv := &http.Server{
		Handler: mux,
	}

	ctx, ctxCancel := context.WithCancel(context.Background())

	app := &Application{
		srv:       srv,
		ctx:       ctx,
		ctxCancel: ctxCancel,
		log:       slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	return app
}

func (a *Application) Run() {
	go a.gracefulShutdown()

	a.srv.Addr = ":" + a.Port

	a.log.Info(
		"application started",
		"port", a.Port,
		"name", version.Name,
		"version", version.Version,
		"commit", version.Commit,
		"build_time", version.BuildTime,
	)

	if err := a.srv.ListenAndServe(); err != nil {
		a.log.Error(err.Error())
	}
}

func (a *Application) gracefulShutdown() {
	ctx := a.ctx

	signalCtx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	<-signalCtx.Done()
	a.log.Info("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(shutdownCtx); err != nil {
		a.log.Error(err.Error())
	}
}
