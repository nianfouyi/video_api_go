package services

import (
	"github.com/nianfouyi/video-user-api/internal/models"
	"github.com/nianfouyi/video-user-api/internal/repositories"
)

type VideoService struct {
	repo *repositories.VideoRepository
}

func NewVideoService(repo *repositories.VideoRepository) *VideoService {
	return &VideoService{repo: repo}
}

func (s *VideoService) CreateVideos(name string) (*models.Video, error) {
	video := &models.Video{Name: name}
	err := s.repo.Create(video)
	return video, err
}

func (s *VideoService) GetAllVideos(filters map[string]interface{}) ([]models.Video, error) {
	return s.repo.FindAll(filters)
}

func (s *VideoService) GetVideoByID(id uint) (*models.Video, error) {
	return s.repo.FindByID(id)
}

func (s *VideoService) UpdateVideo(video *models.Video) error {
	return s.repo.Update(video)
}

func (s *VideoService) DeleteVideo(id uint) error {
	return s.repo.Delete(id)
}
