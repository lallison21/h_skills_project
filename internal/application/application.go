package application

import (
	"context"
	"fmt"
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

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("/status"))
		if err != nil {
			fmt.Println("error to write")
		}
	})

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
