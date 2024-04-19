package model

import "time"

type Comment struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	TimelineId uint      `json:"timeline_id" gorm:"not null"`
	UserId     uint      `json:"user_id" gorm:"not null"`
	Comment    string    `json:"comment" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}