package repositories

import (
	"github.com/nianfouyi/video-user-api/internal/models"
	"gorm.io/gorm"
)

type TagRepository struct {
	DB *gorm.DB
}

type TagRepositoryInterface interface {
	Create(tag *models.Tag) error
	FindByID(id uint) (*models.Tag, error)
	FindByUsername(username string) (*models.Tag, error)
	Update(tag *models.Tag) error
	Delete(id uint) error
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{DB: db}
}

func (r *TagRepository) Create(tag *models.Tag) error {
	return r.DB.Create(tag).Error
}

func (r *TagRepository) FindByID(id uint) (*models.Tag, error) {
	var tag models.Tag
	err := r.DB.First(&tag, id).Error
	return &tag, err
}

func (r *TagRepository) FindAll() ([]models.Tag, error) {
	var tags []models.Tag
	err := r.DB.Find(&tags).Error
	return tags, err
}

func (r *TagRepository) Update(tag *models.Tag) error {
	return r.DB.Save(tag).Error
}

func (r *TagRepository) Delete(id uint) error {
	// 使用事务来确保所有操作都成功完成或全部回滚
	return r.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 获取要删除的tag
		var tag models.Tag
		if err := tx.First(&tag, id).Error; err != nil {
			// 如果找不到tag或发生其他错误，返回错误
			return err
		}

		// 2. 从所有相关视频中移除这个tag
		if err := tx.Model(&tag).Association("Videos").Clear(); err != nil {
			// 如果清除关联失败，返回错误
			return err
		}

		// 3. 删除tag
		if err := tx.Delete(&tag).Error; err != nil {
			// 如果删除tag失败，返回错误
			return err
		}

		// 如果所有操作都成功，事务将被提交
		return nil
	})
}
