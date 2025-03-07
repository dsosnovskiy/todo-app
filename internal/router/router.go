package router

import (
	"todo-app/internal/handler"

	"github.com/go-chi/chi/v5"
)

func NewRouter(taskHandler *handler.TaskHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/tasks", taskHandler.GetAllTasks)
	r.Get("/tasks/{id}", taskHandler.GetTaskByID)
	r.Post("/tasks", taskHandler.CreateTask)
	r.Put("/tasks/{id}", taskHandler.UpdateTask)
	r.Delete("/tasks/{id}", taskHandler.DeleteTask)

	return r
}
