package repository

import (
	"gorm.io/gorm"
	"guthub.com/Junx27/event-app/entity"
)

type UserReopository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserReopository {
	return &UserReopository{db: db}
}

func (ur *UserReopository) GetAll() ([]*entity.User, error) {
	var users []*entity.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
