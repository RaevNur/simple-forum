package main

import (
	"forum/internal/database"
	"forum/internal/helper"
	"forum/internal/repository"
	"forum/internal/service"
	"forum/internal/web"
	"log"
	"os"
)

func main() {
	// set the envoirments
	helper.SetEnv()

	// open the db
	db, err := database.InitDB()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// create repos and services
	repos := repository.NewRepo(db)
	services := service.NewService(repos)

	// create handlers
	handlers, err := web.NewMainHandler(services)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	server := new(web.Server)
	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
