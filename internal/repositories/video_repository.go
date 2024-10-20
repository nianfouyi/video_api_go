package repositories

import (
	"github.com/nianfouyi/video-user-api/internal/models"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type VideoRepository struct {
	DB *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepository {
	return &VideoRepository{DB: db}
}

func (r *VideoRepository) Create(video *models.Video) error {
	return r.DB.Create(video).Error
}

func (r *VideoRepository) FindAll(filters map[string]interface{}) ([]models.Video, error) {
	var videos []models.Video
	query := r.DB.Preload("Tags")

	for key, value := range filters {
		query = query.Where(key+" LIKE ?", "%"+value.(string)+"%")
	}

	err := query.Find(&videos).Error
	return videos, err
}

func (r *VideoRepository) FindByID(id uint) (*models.Video, error) {
	var video models.Video
	err := r.DB.Preload("Tags").First(&video, id).Error
	return &video, err
}

func (r *VideoRepository) Update(video *models.Video) error {
	return r.DB.Save(video).Error
}

func (r *VideoRepository) Delete(id uint) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 获取要删除的视频
		var video models.Video
		if err := tx.First(&video, id).Error; err != nil {
			return err
		}

		// 2. 删除与标签的关联
		if err := tx.Model(&video).Association("Tags").Clear(); err != nil {
			return err
		}

		// 3. 删除视频文件（如果有）
		if video.FilePath != "" {
			if err := os.Remove(filepath.Clean(video.FilePath)); err != nil {
				// 如果文件不存在，我们可能想继续删除数据库记录
				if !os.IsNotExist(err) {
					return err
				}
			}
		}

		// 4. 删除视频记录
		if err := tx.Delete(&video).Error; err != nil {
			return err
		}

		return nil
	})
}
