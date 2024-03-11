package repository

import (
	"github.com/BerkatPS/goRPC-chat/internal/app/model"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (u *userRepositoryImpl) FindByUsername(username string) (*model.User, error) {
	var user model.User

	err := u.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (u *userRepositoryImpl) Save(user *model.User) error {
	return u.DB.Create(user).Error
}
