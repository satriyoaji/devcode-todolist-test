package models

import "time"

type Todo struct {
	ID              int       `json:"id"`
	ActivityGroupID int       `json:"activity_group_id"`
	Title           string    `json:"title"`
	Priority        string    `json:"priority"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
