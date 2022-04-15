package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"todo-api/config"
	"todo-api/model"
)

func DBConnect(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(cfg.Database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})

	return db
}
