package main

import (
	"database/sql"
	computer_gateways "github.com/lallison/h_skills_project/internal/service/gateways/computer"
	computer_repository "github.com/lallison/h_skills_project/internal/service/repository/computer"
	computer_usecase "github.com/lallison/h_skills_project/internal/service/usecase/computer"
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

	_ = computerGateways

	app.Run()
}
