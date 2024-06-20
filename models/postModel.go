package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) (err error) {
	post.ID = time.Now().UnixMilli()
	return
}
