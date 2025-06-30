package models

import "time"

const (
	StatusInProgress = "in_progress"
	StatusCompleted  = "completed"
	StatusFailed     = "failed"
)

// Task описывает полную модель задачи
// @Description Полная модель задачи
type Task struct {
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Created_at   time.Time  `json:"created_at" example:"2023-05-15T10:00:00Z" format:"date-time"`
	Completed_at *time.Time `json:"completed_at,omitempty" example:"2023-05-15T10:00:00Z" format:"date-time"`
	Duration     *int       `json:"duration,omitempty"` // в минутах
	Result       *string    `json:"result,omitempty"`
	Status       string     `json:"status"`
}

// CreateTaskRequst описывает упрощенную модель задачи для добавления
// @Description Модель задачи для запроса на создание
type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
