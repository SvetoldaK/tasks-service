package grpc

import (
	"net"

	"google.golang.org/grpc"

	tasksService "tasks-service/internal/task"

	taskpb "github.com/SvetoldaK/project-protos/proto/task"
	userpb "github.com/SvetoldaK/project-protos/proto/user"
)

func RunGRPC(svc *tasksService.TaskService, uc userpb.UserServiceClient) error {
	lis, _ := net.Listen("tcp", ":50052")
	grpcSrv := grpc.NewServer()
	handler := NewHandler(svc, uc)
	taskpb.RegisterTaskServiceServer(grpcSrv, handler)
	return grpcSrv.Serve(lis)
}
