package persistence

import (
	"fmt"
	"gorm.io/gorm"
	"todo-api/config"
	"todo-api/model"
)

type User interface {
	FindByEmail(email string) (*model.User, error)
}

type userPesistence struct {
	db  *gorm.DB
	cfg *config.Config
}

func (u userPesistence) FindByEmail(email string) (*model.User, error) {
	var user *model.User
	result := u.db.Table(u.cfg.DatabaseUserCollection).Find(&user, "email = ?", email)
	if result.Error != nil {
		return nil, fmt.Errorf("error in found user", result.Error.Error())
	}
	return user, nil
}

func NewPersistenceUser(db *gorm.DB, cfg *config.Config) User {
	return &userPesistence{
		db:  db,
		cfg: cfg,
	}
}
