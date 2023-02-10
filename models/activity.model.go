package models

import "time"

type Activity struct {
	ID        int        `json:"id" gorm:"primaryKey,autoIncrement"`
	Title     string     `json:"title"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
