package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/SemenTretyakov/chat_server/pkg/chat_v1"
	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const PORT = 50052

type server struct {
	desc.UnimplementedChatV1Server
}

func (s *server) Create(ctx context.Context, req *desc.CreateReq) (*desc.CreateRes, error) {
	if req.GetUsernames() == nil {
		log.Fatalf("Not provide user names for create chat")
	}

	log.Printf("Creating chat: %v",
		req.GetUsernames(),
	)

	return &desc.CreateRes{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageReq) (*empty.Empty, error) {
	if req.GetFrom() == "" || req.GetText() == "" {
		log.Fatalf("Not provide user names for create chat")
	}

	log.Printf("Send message with text: %v and from:", req.GetText(), req.GetFrom())
	return &empty.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteReq) (*empty.Empty, error) {
	if req.GetId() == 0 {
		return &empty.Empty{}, status.Error(codes.InvalidArgument, "Chat ID is required")
	}

	log.Printf("Deleting Chat with id: %d", req.GetId())
	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to Serve: %v", err)
	}
}
