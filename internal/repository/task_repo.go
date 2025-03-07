package repository

import (
	"todo-app/internal/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetAllTasks() ([]model.Task, error)
	GetTaskByID(taskID uint) (*model.Task, error)
	CreateTask(task *model.Task) error
	UpdateStatusTask(taskID uint, status bool) error
	UpdateTitleTask(taskID uint, title string) error
	DeleteTask(taskID uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTaskByID(taskID uint) (*model.Task, error) {
	var task model.Task
	if err := r.db.Model(&model.Task{}).Where("task_id = ?", taskID).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) CreateTask(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) UpdateStatusTask(taskID uint, status bool) error {
	var task model.Task
	if err := r.db.Model(&model.Task{}).Where("task_id = ?", taskID).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&task).Update("status", status).Error
}

func (r *taskRepository) UpdateTitleTask(taskID uint, title string) error {
	var task model.Task
	if err := r.db.Model(&model.Task{}).Where("task_id = ?", taskID).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&task).Update("title", title).Error
}

func (r *taskRepository) DeleteTask(taskID uint) error {
	var task model.Task
	if err := r.db.Model(&model.Task{}).Where("task_id = ?", taskID).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Delete(&task).Error
}
