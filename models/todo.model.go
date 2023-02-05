package models

import "time"

type Todo struct {
	TodoID     int       `json:"todo_id"`
	ActivityID int       `json:"activity_id"`
	Title      string    `json:"title" validate:"required"`
	Priority   string    `json:"priority" validate:"required"`
	IsActive   bool      `json:"is_active" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
