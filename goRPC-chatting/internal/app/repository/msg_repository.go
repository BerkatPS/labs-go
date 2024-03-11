package repository

import "github.com/BerkatPS/goRPC-chat/internal/app/model"

type MessageRepository interface {
	Save(message *model.Message) error
	FindByRecipient(username string) ([]*model.Message, error)
}
