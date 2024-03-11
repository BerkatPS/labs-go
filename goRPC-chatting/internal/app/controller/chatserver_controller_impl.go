package controller

import (
	"context"
	"github.com/BerkatPS/goRPC-chat/internal/app/model"
	labs_go "github.com/BerkatPS/goRPC-chat/internal/app/proto"
	"github.com/BerkatPS/goRPC-chat/internal/app/service"
)

type ChatServerImpl struct {
	userService    service.UserService
	messageService service.MessageService
}

// NewChatServer creates a new instance of ChatServerImpl.
func NewChatServer(userService service.UserService, messageService service.MessageService) *ChatServerImpl {
	return &ChatServerImpl{
		userService:    userService,
		messageService: messageService,
	}
}

// SendMessage implements the SendMessage RPC method of the ChatServiceServer interface.
func (c *ChatServerImpl) SendMessage(ctx context.Context, req *labs_go.MessageRequest) (*labs_go.MessageResponse, error) {
	// Find sender user
	sender, err := c.userService.FindByUsername(req.Sender)
	if err != nil {
		return nil, err
	}

	// Find recipient user
	recipient, err := c.userService.FindByUsername(req.Recipient)
	if err != nil {
		return nil, err
	}

	// Save message to database
	message := &model.Message{
		Sender:    sender.Username,
		Recipient: recipient.Username,
		Content:   req.Content,
	}

	if err := c.messageService.Save(message); err != nil {
		return nil, err
	}

	return &labs_go.MessageResponse{Status: "Message Sent Successfully"}, nil
}

// GetMessageForUser implements the GetMessageForUser RPC method of the ChatServiceServer interface.
func (c *ChatServerImpl) GetMessageForUser(ctx context.Context, req *labs_go.UserRequest) (*labs_go.MessagesResponse, error) {
	// Find messages for user
	messages, err := c.messageService.FindByRecipient(req.Username)
	if err != nil {
		return nil, err
	}

	// Format messages into response
	var response []*labs_go.MessageData
	for _, msg := range messages {
		response = append(response, &labs_go.MessageData{
			Sender:    msg.Sender,
			Recipient: msg.Recipient,
			Content:   msg.Content,
		})
	}

	return &labs_go.MessagesResponse{Messages: response}, nil
}
func (c *ChatServerImpl) mustEmbedUnimplementedChatServiceServer() {

}
