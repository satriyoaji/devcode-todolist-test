package dto

import "time"

type Todo struct {
	ID              int       `json:"id"`
	ActivityGroupID int       `json:"activity_group_id"`
	Title           string    `json:"title" validate:"required"`
	Priority        string    `json:"priority" validate:"required"`
	IsActive        bool      `json:"is_active" validate:"required"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
