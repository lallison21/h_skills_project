package main

import (
	"fmt"
	"os"

	"github.com/lallison/h_skills_project/internal/application"
)

func main() {
	port := os.Getenv("APPLICATION_PORT")
	if port == "" {
		fmt.Printf("enviroment variable APPLICATION_PORT is not set")
		return
	}

	// db, err := sqlx.Connect()
	// if err != nil {
	// 	// ...
	// }

	// computerRepository := computer_repo.New(db)
	// computerUseCase := computer_uc.New(computerRepository)

	app := application.New()
	app.Port = port

	app.Run()
}
