package service

import (
	"fmt"
	"todo-app/internal/model"
	"todo-app/internal/repository"
)

type TaskService struct {
	taskRepo *repository.TaskRepository
}

func NewTaskService(taskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	tasks, err := s.taskRepo.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tasks: %v", err)
	}
	return tasks, nil
}

func (s *TaskService) GetTaskByID(taskID uint) (*model.Task, error) {
	task, err := s.taskRepo.GetTaskByID(taskID)
	if err != nil {
		return nil, fmt.Errorf("task not found: %v", err)
	}
	return task, nil
}

func (s *TaskService) CreateTask(task *model.Task) error {
	if task.Title == "" {
		return fmt.Errorf("task title cannot be empty")
	}
	if err := s.taskRepo.CreateTask(task); err != nil {
		return fmt.Errorf("failed to create task: %v", err)
	}
	return nil
}

func (s *TaskService) UpdateStatusTask(taskID uint, status bool) error {
	if err := s.taskRepo.UpdateStatusTask(taskID, status); err != nil {
		return fmt.Errorf("failed to update task status: %v", err)
	}
	return nil
}

func (s *TaskService) UpdateTitleTask(taskID uint, title string) error {
	if title == "" {
		return fmt.Errorf("task title cannot be empty")
	}
	if err := s.taskRepo.UpdateTitleTask(taskID, title); err != nil {
		return fmt.Errorf("failed to update task title: %v", err)
	}
	return nil
}

func (s *TaskService) DeleteTask(taskID uint) error {
	if err := s.taskRepo.DeleteTask(taskID); err != nil {
		return fmt.Errorf("failed to delete task: %v", err)
	}
	return nil
}
