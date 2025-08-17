package main

import (
	"github.com/filtkac/consult-scheduler/config"
	"log"
)

func main() {
	db := config.DatabaseConnection()
	config.AutoMigrate(db)

	router := config.RouterConfig(db)
	err := router.Run()
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
