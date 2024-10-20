package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username         string    `gorm:"uniqueIndex;not null" json:"username"`
	Password         string    `json:"-"`
	Gender           string    `json:"gender"`
	Hobbies          string    `json:"hobbies"`
	RegisterDate     time.Time `json:"register_date"`
	SecurityQuestion string    `json:"security_question"`
	SecurityAnswer   string    `json:"-"`
}
