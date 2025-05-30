package main

import (
	"log"
	"tasks-service/internal/database"

	tasksService "tasks-service/internal/task"
	transportgrpc "tasks-service/internal/transport/grpc"
)

func main() {
	// 1. Инициализация БД
	database.InitDB()
	database.DB.AutoMigrate(&tasksService.Task{})

	// 2. Репозиторий и сервис задач
	repo := tasksService.NewTaskRepository(database.DB)
	svc := tasksService.NewService(repo)

	// 3. Клиент к Users-сервису
	userClient, conn, err := transportgrpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect to users: %v", err)
	}
	defer conn.Close()

	// 4. Запуск gRPC Tasks-сервиса
	if err := transportgrpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}
