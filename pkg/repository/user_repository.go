package repository

import (
	"github.com/tinwritescode/gin-tin/pkg/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}
