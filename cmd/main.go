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

	cfg := config.Load()

	logger := logging.Init(cfg.Env)

	db, err := database.Connect(cfg)
	if err != nil {

		logger.Fatalf("Failed to connect to the database: %v", err)
	}
	logger.Debug("Successfully connected to the database")

	taskRepo := repository.NewTaskRepository(db)
	logger.Debug("Successfully initialized the repository")

	taskService := service.NewTaskService(taskRepo)
	logger.Debug("Successfully initialized the service")

	taskHandler := handler.NewTaskHandler(taskService)
	logger.Debug("Successfully initialized the handler")

	r := router.New(taskHandler)
	logger.Debug("Successfully initialized the router")

	logger.Infof("The server is running on port: %s", cfg.Server.Port)
	if err := http.ListenAndServe("localhost:"+cfg.Server.Port, r); err != nil {
		logger.Fatalf("Failed to start the server: %v", err)
	}
}
