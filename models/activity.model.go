package models

import "time"

type Activity struct {
	ActivityID int       `json:"activity_id"`
	Title      string    `json:"title" validate:"required"`
	Email      string    `json:"email" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
