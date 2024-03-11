package main

import (
	"log"
	"net"

	"github.com/BerkatPS/goRPC-chat/config"
	"github.com/BerkatPS/goRPC-chat/internal/app/controller"
	labs_go "github.com/BerkatPS/goRPC-chat/internal/app/proto"
	"github.com/BerkatPS/goRPC-chat/internal/app/repository"
	"github.com/BerkatPS/goRPC-chat/internal/app/service"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func main() {
	// Initialize database configuration
	dbConfig := config.NewConfig()
	db := dbConfig.DB

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
	})

	// Initialize repositories
	userRepository := repository.NewUserRepository(db)
	messageRepository := repository.NewMessageRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepository)
	messageService := service.NewMessageService(messageRepository)

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	chatController := controller.NewChatServer(userService, messageService)

	var chatServiceServer labs_go.ChatServiceServer = chatController

	labs_go.RegisterChatServiceServer(grpcServer, chatServiceServer)

	// Start gRPC server
	go func() {
		listen, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Println("gRPC server is running on port 50051")
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Example endpoint for REST API
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Testing endpoint")
	})

	// Run Fiber web server
	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("failed to start Fiber server: %v", err)
	}
}
