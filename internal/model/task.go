package model

import "time"

type Task struct {
	TaskID    uint      `gorm:"primaryKey;autoIncrement" json:"task_id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Status    bool      `gorm:"default:false" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
