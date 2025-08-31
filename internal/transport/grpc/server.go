package grpc

import (
	"fmt"
	"net"

	taskpb "github.com/Vishnevyy/project-protos/proto/task"
	userpb "github.com/Vishnevyy/project-protos/proto/user"

	domain "github.com/Vishnevyy/tasks-service/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc *domain.Service, uc userpb.UserServiceClient) error {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return fmt.Errorf("failed to listen on :50052: %w", err)
	}

	s := grpc.NewServer()
	taskpb.RegisterTaskServiceServer(s, NewHandler(svc, uc))

	fmt.Println("Tasks gRPC server listening on :50052")
	return s.Serve(lis)
}
