package models

import "time"

type Todo struct {
	ID              int        `json:"id" gorm:"primaryKey,autoIncrement"`
	ActivityGroupID int        `json:"activity_group_id"`
	Title           string     `json:"title"`
	Priority        string     `json:"priority"`
	IsActive        bool       `json:"is_active"`
	CreatedAt       *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
