package services

import (
	"github.com/nianfouyi/video-user-api/internal/models"
	"github.com/nianfouyi/video-user-api/internal/repositories"
	"gorm.io/gorm"
)

type PlaybackRecordService interface {
	RecordPlayback(userID, videoID uint, watchTime int) error
}

type playbackRecordService struct {
	repo repositories.PlaybackRecordRepository
}

func NewPlaybackRecordService(repo repositories.PlaybackRecordRepository) PlaybackRecordService {
	return &playbackRecordService{repo: repo}
}

func (s *playbackRecordService) RecordPlayback(userID, videoID uint, watchTime int) error {
	record, err := s.repo.FindByUserAndVideo(userID, videoID)
	if err != nil {
		// 如果记录不存在，创建新记录
		if err == gorm.ErrRecordNotFound {
			newRecord := &models.PlaybackRecord{
				UserID:    userID,
				VideoID:   videoID,
				WatchTime: watchTime,
			}
			return s.repo.Create(newRecord)
		}
		return err
	}

	// 更新现有记录
	record.WatchTime += watchTime
	return s.repo.Update(record)
}
