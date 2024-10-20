package services

import (
	"github.com/nianfouyi/video-user-api/internal/models"
	"github.com/nianfouyi/video-user-api/internal/repositories"
)

type TagService struct {
	repo *repositories.TagRepository
}

func NewTagService(repo *repositories.TagRepository) *TagService {
	return &TagService{repo: repo}
}

func (s *TagService) CreateTag(name string) (*models.Tag, error) {
	tag := &models.Tag{Name: name}
	err := s.repo.Create(tag)
	return tag, err
}

func (s *TagService) GetTagByID(id uint) (*models.Tag, error) {
	return s.repo.FindByID(id)
}

func (s *TagService) GetAllTags() ([]models.Tag, error) {
	return s.repo.FindAll()
}

func (s *TagService) UpdateTag(tag *models.Tag) error {
	return s.repo.Update(tag)
}

func (s *TagService) DeleteTag(id uint) error {
	return s.repo.Delete(id)
}
