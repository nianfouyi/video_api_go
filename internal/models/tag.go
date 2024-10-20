package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name   string  `gorm:"uniqueIndex;not null" json:"name"`
	Videos []Video `gorm:"many2many:video_tags;" json:"videos,omitempty"`
}
