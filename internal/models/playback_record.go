package models

import (
	"time"
)

type PlaybackRecord struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	VideoID   uint      `json:"video_id"`
	WatchTime int       `json:"watch_time"` // 观看时长（秒）
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
