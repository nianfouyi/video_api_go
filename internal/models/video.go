package models

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name      string    `json:"name"`
	Number    string    `json:"number"`
	Category  string    `json:"category"`
	MainActor string    `json:"main_actor"`
	Rating    float32   `json:"rating"`
	AddedAt   time.Time `json:"added_at"`
	Tags      []Tag     `gorm:"many2many:video_tags;" json:"tags"`
	FilePath  string
}
