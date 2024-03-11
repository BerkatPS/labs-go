package repository

import "github.com/BerkatPS/goRPC-chat/internal/app/model"

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	Save(user *model.User) error
}
