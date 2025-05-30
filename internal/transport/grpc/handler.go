package grpc

import (
	"context"
	"fmt"

	taskpb "github.com/SvetoldaK/project-protos/proto/task"
	userpb "github.com/SvetoldaK/project-protos/proto/user"

	tasksService "tasks-service/internal/task"
)

type Handler struct {
	svc        *tasksService.TaskService
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc *tasksService.TaskService, uc userpb.UserServiceClient) *Handler {
	return &Handler{svc: svc, userClient: uc}
}

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	t, err := h.svc.CreateTask(tasksService.Task{
		Task: req.Title,
	})
	if err != nil {
		return nil, err
	}
	return &taskpb.CreateTaskResponse{
		Task: &taskpb.Task{
			Id:    uint32(t.ID),
			Title: t.Task,
		},
	}, nil
}

func (h *Handler) GetTask(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {
	tasks, err := h.svc.GetAllTasks()
	if err != nil {
		return nil, err
	}
	for _, t := range tasks {
		if uint32(t.ID) == req.Id {
			return &taskpb.GetTaskResponse{
				Task: &taskpb.Task{
					Id:    uint32(t.ID),
					Title: t.Task,
				},
			}, nil
		}
	}
	return nil, fmt.Errorf("task with id %d not found", req.Id)
}

func (h *Handler) ListTasks(ctx context.Context, req *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.GetAllTasks()
	if err != nil {
		return nil, err
	}
	var pbTasks []*taskpb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &taskpb.Task{
			Id:    uint32(t.ID),
			Title: t.Task,
		})
	}
	return &taskpb.ListTasksResponse{
		Tasks:      pbTasks,
		TotalCount: int32(len(pbTasks)),
	}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	t, err := h.svc.UpdateTaskByID(uint(req.Id), tasksService.Task{
		Task: req.Title,
	})
	if err != nil {
		return nil, err
	}
	return &taskpb.UpdateTaskResponse{
		Task: &taskpb.Task{
			Id:    uint32(t.ID),
			Title: t.Task,
		},
	}, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	err := h.svc.DeleteTaskByID(uint(req.Id))
	if err != nil {
		return &taskpb.DeleteTaskResponse{Success: false}, err
	}
	return &taskpb.DeleteTaskResponse{Success: true}, nil
}
