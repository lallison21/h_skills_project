package application

import (
	"context"
	"fmt"
	gateway_status "github.com/lallison/h_skills_project/internal/service/gateways/status"
	repository_status "github.com/lallison/h_skills_project/internal/service/repository/status"
	usecase_status "github.com/lallison/h_skills_project/internal/service/usecase/status"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server interface {
	Run()
	Shutdown(ctx context.Context) error
}

type application struct {
	srv *http.Server
}

func New() Server {
	const fn = "application.New"

	addr := os.Getenv("APPLICATION_PORT")
	if addr == "" {
		log.Fatalf("%s: enviroment variable APPLICATION_PORT is not set", fn)
	}

	mux := http.NewServeMux()

	repository := repository_status.New()
	useCase := usecase_status.New(repository)
	gateways := gateway_status.New(useCase)

	mux.HandleFunc("/status", gateways.HandleGetStatus)

	return &application{
		srv: &http.Server{
			Addr:    fmt.Sprintf("0.0.0.0:%s", addr),
			Handler: mux,
		},
	}
}

func (a *application) Run() {
	const fn = "application.Run"

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	_ = ctx

	if err := a.srv.ListenAndServe(); err != nil {
		log.Fatalf("%s: can't ruuning server: %s", fn, err.Error())
	}
}

func (a *application) Shutdown(ctx context.Context) error {
	return a.srv.Shutdown(ctx)
}
