package service

import (
	"github.com/BerkatPS/goRPC-chat/internal/app/model"
	"github.com/BerkatPS/goRPC-chat/internal/app/repository"
)

type messageServiceImpl struct {
	msgRepo repository.MessageRepository
}

func NewMessageService(msgRepo repository.MessageRepository) MessageService {
	return &messageServiceImpl{msgRepo: msgRepo}
}

func (m *messageServiceImpl) Save(message *model.Message) error {
	return m.msgRepo.Save(message)
}

func (m *messageServiceImpl) FindByRecipient(username string) ([]*model.Message, error) {
	return m.msgRepo.FindByRecipient(username)
}
