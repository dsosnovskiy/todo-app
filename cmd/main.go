package main

import (
	"net/http"
	"todo-app/config"
	"todo-app/internal/handler"
	"todo-app/internal/repository"
	"todo-app/internal/router"
	"todo-app/internal/service"
	"todo-app/pkg/database"
	"todo-app/pkg/logging"
)

func main() {

	// Load the configuration
	cfg := config.LoadConfig()

	// Initialize the logger
	log := logging.InitLogger(cfg.Environment)

	// Creating models
	// Connection to the database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Debug("Successfully connected to the database")

	// Initialize the repository
	taskRepo := repository.NewTaskRepository(db)
	log.Debug("Successfully initialized the repository")

	// Initialize the service
	taskService := service.NewTaskService(taskRepo)
	log.Debug("Successfully initialized the service")

	// Initialize the handler
	taskHandler := handler.NewTaskHandler(taskService)
	log.Debug("Successfully initialized the handler")

	// Initialize the router
	r := router.NewRouter(taskHandler)
	log.Debug("Successfully initialized the router")

	// Start the server
	log.Debug("The server is running on port: ", cfg.Server.Port)
	http.ListenAndServe(":"+cfg.Server.Port, r)
}
