package service

import (
	"errors"
	"fmt"
	"todo-app/internal/model"
	"todo-app/internal/repository"
)

var (
	ErrTaskNotFound  = errors.New("task not found")
	ErrEmptyTitle    = errors.New("task title cannot be empty")
	ErrRetrieveTasks = errors.New("failed to retrieve tasks")
	ErrCreateTask    = errors.New("failed to create task")
	ErrUpdateStatus  = errors.New("failed to update task status")
	ErrUpdateTitle   = errors.New("failed to update task title")
	ErrDeleteTask    = errors.New("failed to delete task")
)

type TaskService interface {
	GetAllTasks() ([]model.Task, error)
	GetTaskByID(taskID uint) (*model.Task, error)
	CreateTask(task *model.Task) error
	UpdateTitleTask(taskID uint, title string) error
	UpdateStatusTask(taskID uint, status bool) error
	DeleteTask(taskID uint) error
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{taskRepo: taskRepo}
}

func (s *taskService) GetAllTasks() ([]model.Task, error) {
	tasks, err := s.taskRepo.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRetrieveTasks, err)
	}
	return tasks, nil
}

func (s *taskService) GetTaskByID(taskID uint) (*model.Task, error) {
	task, err := s.taskRepo.GetTaskByID(taskID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrTaskNotFound, err)
	}
	return task, nil
}

func (s *taskService) CreateTask(task *model.Task) error {
	if task.Title == "" {
		return ErrEmptyTitle
	}
	if err := s.taskRepo.CreateTask(task); err != nil {
		return fmt.Errorf("%w: %v", ErrCreateTask, err)
	}
	return nil
}

func (s *taskService) UpdateStatusTask(taskID uint, status bool) error {
	if err := s.taskRepo.UpdateStatusTask(taskID, status); err != nil {
		return fmt.Errorf("%w: %v", ErrUpdateStatus, err)
	}
	return nil
}

func (s *taskService) UpdateTitleTask(taskID uint, title string) error {
	if title == "" {
		return ErrEmptyTitle
	}
	if err := s.taskRepo.UpdateTitleTask(taskID, title); err != nil {
		return fmt.Errorf("%w: %v", ErrUpdateTitle, err)
	}
	return nil
}

func (s *taskService) DeleteTask(taskID uint) error {
	if err := s.taskRepo.DeleteTask(taskID); err != nil {
		return fmt.Errorf("%w: %v", ErrDeleteTask, err)
	}
	return nil
}
