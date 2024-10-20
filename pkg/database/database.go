package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// 使用 SQLite 数据库文件
	dbPath := "videouserdb.sqlite"

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 在这里进行数据库迁移
	// db.AutoMigrate(&models.User{}, &models.Video{}, &models.Tag{})

	return db, nil
}
