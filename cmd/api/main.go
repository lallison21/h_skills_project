package main

import (
	"database/sql"
	"github.com/lallison/h_skills_project/internal/service/gateways/computer"
	"github.com/lallison/h_skills_project/internal/service/gateways/status"
	"github.com/lallison/h_skills_project/internal/service/repository/computer"
	"github.com/lallison/h_skills_project/internal/service/repository/status"
	"github.com/lallison/h_skills_project/internal/service/usecase/computer"
	"github.com/lallison/h_skills_project/internal/service/usecase/status"
	_ "github.com/lib/pq"
	"log"
	"os"

	"github.com/lallison/h_skills_project/internal/application"
)

func main() {
	addr := os.Getenv("APPLICATION_ADDR")
	if addr == "" {
		log.Fatal("environment variable APPLICATION_PORT is not set")
		return
	}

	connStr := os.Getenv("APPLICATION_DB")
	if connStr == "" {
		log.Fatal("environment variable APPLICATION_DB is not set")
		return
	}

	app := application.New()
	app.Addr = addr

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	computerRepo := computer_repository.New(db)
	computerUseCase := computer_usecase.New(computerRepo)
	computerGateways := computer_gateways.New(computerUseCase)
	app.RegisterComputerGateways(computerGateways)

	statusRepository := repository_status.New(db)
	statusUseCase := usecase_status.New(statusRepository)
	statusGateways := gateway_status.New(statusUseCase)
	app.RegisterStatusGateways(statusGateways)

	app.Run()

	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
