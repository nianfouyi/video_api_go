package repositories

import (
	"github.com/nianfouyi/video-user-api/internal/models"
	"gorm.io/gorm"
)

type playbackRecordRepository struct {
	db *gorm.DB
}

type PlaybackRecordRepository interface {
	Create(record *models.PlaybackRecord) error
	FindByUserAndVideo(userID, videoID uint) (*models.PlaybackRecord, error)
	Update(record *models.PlaybackRecord) error
}

func NewPlaybackRecordRepository(db *gorm.DB) PlaybackRecordRepository {
	return &playbackRecordRepository{db: db}
}

func (r *playbackRecordRepository) Create(record *models.PlaybackRecord) error {
	return r.db.Create(record).Error
}

func (r *playbackRecordRepository) FindByUserAndVideo(userID, videoID uint) (*models.PlaybackRecord, error) {
	var record models.PlaybackRecord
	err := r.db.Where("user_id = ? AND video_id = ?", userID, videoID).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *playbackRecordRepository) Update(record *models.PlaybackRecord) error {
	return r.db.Save(record).Error
}
