package service

import "github.com/BerkatPS/goRPC-chat/internal/app/model"

type UserService interface {
	FindByUsername(username string) (*model.User, error)
	Save(user *model.User) error
}
