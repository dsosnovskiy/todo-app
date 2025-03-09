package repository

import (
	"todo-app/internal/model"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetTaskByID(taskID uint) (*model.Task, error) {
	var task model.Task
	if err := r.db.Model(&model.Task{}).Where("task_id = ?", taskID).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) CreateTask(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) UpdateStatusTask(taskID uint, status bool) error {
	var task model.Task
	if err := r.db.Model(&model.Task{}).Where("task_id = ?", taskID).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&task).Update("status", status).Error
}

func (r *TaskRepository) UpdateTitleTask(taskID uint, title string) error {
	var task model.Task
	if err := r.db.Model(&model.Task{}).Where("task_id = ?", taskID).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&task).Update("title", title).Error
}

func (r *TaskRepository) DeleteTask(taskID uint) error {
	var task model.Task
	if err := r.db.Model(&model.Task{}).Where("task_id = ?", taskID).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Delete(&task).Error
}
