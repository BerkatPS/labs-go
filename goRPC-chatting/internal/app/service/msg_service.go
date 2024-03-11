package service

import "github.com/BerkatPS/goRPC-chat/internal/app/model"

type MessageService interface {
	Save(message *model.Message) error
	FindByRecipient(username string) ([]*model.Message, error)
}
