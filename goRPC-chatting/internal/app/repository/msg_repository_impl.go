package repository

import (
	"github.com/BerkatPS/goRPC-chat/internal/app/model"
	"gorm.io/gorm"
)

type messageRepositoryImpl struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepositoryImpl{DB: db}
}
func (m *messageRepositoryImpl) Save(message *model.Message) error {
	return m.DB.Create(message).Error
}

func (m *messageRepositoryImpl) FindByRecipient(username string) ([]*model.Message, error) {
	var messages []*model.Message
	err := m.DB.Where("recipient = ?", username).Find(&messages).Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}
