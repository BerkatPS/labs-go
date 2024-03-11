package service

import (
	"github.com/BerkatPS/goRPC-chat/internal/app/model"
	"github.com/BerkatPS/goRPC-chat/internal/app/repository"
)

type userServiceImpl struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new instance of UserServiceImpl
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{userRepo: userRepo}
}

// Save saves the user to the database
func (s *userServiceImpl) Save(user *model.User) error {
	return s.userRepo.Save(user)
}

// FindByUsername finds a user by username
func (s *userServiceImpl) FindByUsername(username string) (*model.User, error) {
	return s.userRepo.FindByUsername(username)
}
