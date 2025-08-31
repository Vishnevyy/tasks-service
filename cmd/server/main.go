package main

import (
	"log"

	"github.com/Vishnevyy/tasks-service/internal/database"
	domain "github.com/Vishnevyy/tasks-service/internal/task"
	transportgrpc "github.com/Vishnevyy/tasks-service/internal/transport/grpc"
)

func main() {
	// 1) init DB
	database.InitDB()

	// 2) repo + service
	repo := domain.NewRepository(database.DB)
	svc := domain.NewService(repo)

	// 3) users gRPC client
	userClient, conn, err := transportgrpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect to users service: %v", err)
	}
	defer conn.Close()

	// 4) run tasks gRPC server
	if err := transportgrpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}
