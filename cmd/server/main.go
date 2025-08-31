package main

import (
	"log"

	"github.com/Vishnevyy/tasks-service/internal/database"
	domain "github.com/Vishnevyy/tasks-service/internal/task"
	transportgrpc "github.com/Vishnevyy/tasks-service/internal/transport/grpc"
)

func main() {
	database.InitDB()

	repo := domain.NewRepository(database.DB)
	svc := domain.NewService(repo)

	userClient, conn, err := transportgrpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect to users service: %v", err)
	}
	defer conn.Close()

	if err := transportgrpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}
