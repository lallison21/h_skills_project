package application

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lallison/h_skills_project/version"
)

type Application struct {
	Addr string
	mux  *http.ServeMux
	srv  *http.Server

	ctx       context.Context
	ctxCancel context.CancelFunc

	log *slog.Logger
}

func New() *Application {
	mux := http.NewServeMux()

	srv := &http.Server{}

	ctx, ctxCancel := context.WithCancel(context.Background())

	app := &Application{
		mux:       mux,
		srv:       srv,
		ctx:       ctx,
		ctxCancel: ctxCancel,
		log:       slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}

	return app
}

func (a *Application) Run() {
	go a.gracefulShutdown()

	a.srv.Addr = a.Addr
	a.srv.Handler = a.mux

	a.log.Info(
		"application started",
		"address", a.Addr,
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

type ComputerGateways interface {
	CreateComputer() http.HandlerFunc
	Computers() http.HandlerFunc
	ComputerByID() http.HandlerFunc
}

func (a *Application) RegisterComputerGateways(gateway ComputerGateways) {
	a.mux.HandleFunc("/create_computer", gateway.CreateComputer())
	a.mux.HandleFunc("/computers", gateway.Computers())
	a.mux.HandleFunc("/computers/{id}", gateway.ComputerByID())
}

type StatusGateways interface {
	StatusHandle(w http.ResponseWriter, r *http.Request)
}

func (a *Application) RegisterStatusGateways(gateway StatusGateways) {
	a.mux.HandleFunc("/status", gateway.StatusHandle)
}
