package main

import (
	"context"
	"github.com/BerkatPS/go-rpc/message"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myMessageServer struct {
	message.UnimplementedClientServer
}

func (m myMessageServer) Create(ctx context.Context, req *message.CreateRequest) (*message.CreateResponse, error) {
	return &message.CreateResponse{
		Pdf: []byte(req.From),
		Doc: []byte("123"),
	}, nil
}
func main() {

	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Cannot Create listener: %s", err)
	}

	serverRegister := grpc.NewServer()
	service := &myMessageServer{}

	message.RegisterClientServer(serverRegister, service)

	err = serverRegister.Serve(listen)
	if err != nil {
		log.Fatalf("impposible to serve: %s", err)
	}
}
