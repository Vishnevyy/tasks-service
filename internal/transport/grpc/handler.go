package grpc

import (
	"context"
	"fmt"

	taskpb "github.com/Vishnevyy/project-protos/proto/task"
	userpb "github.com/Vishnevyy/project-protos/proto/user"
	"github.com/Vishnevyy/tasks-service/internal/task"
)

type Handler struct {
	svc        *task.Service
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc *task.Service, uc userpb.UserServiceClient) *Handler {
	return &Handler{
		svc:        svc,
		userClient: uc,
	}
}

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	// Проверяем пользователя
	if _, err := h.userClient.GetUser(ctx, &userpb.User{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}

	// Создаем задачу
	t, err := h.svc.CreateTask(task.Task{
		UserID: req.UserId,
		Title:  req.Title,
	})
	if err != nil {
		return nil, err
	}

	return &taskpb.CreateTaskResponse{
		Task: &taskpb.Task{
			Id:     t.ID,
			UserId: t.UserID,
			Title:  t.Title,
			IsDone: t.IsDone,
		},
	}, nil
}

func (h *Handler) GetTask(ctx context.Context, req *taskpb.Task) (*taskpb.Task, error) {
	t, err := h.svc.GetTask(req.Id)
	if err != nil {
		return nil, err
	}
	return &taskpb.Task{
		Id:     t.ID,
		UserId: t.UserID,
		Title:  t.Title,
		IsDone: t.IsDone,
	}, nil
}

func (h *Handler) ListTasks(ctx context.Context, req *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.ListTasks()
	if err != nil {
		return nil, err
	}

	items := make([]*taskpb.Task, 0, len(tasks))
	for _, t := range tasks {
		items = append(items, &taskpb.Task{
			Id:     t.ID,
			UserId: t.UserID,
			Title:  t.Title,
			IsDone: t.IsDone,
		})
	}

	return &taskpb.ListTasksResponse{Items: items}, nil
}

// Нужно добавить метод ListTasksByUser если он есть в контракте
// Если нет - можно использовать ListTasks с фильтрацией по userId

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.Task, error) {
	// Получаем текущую задачу чтобы узнать userId
	currentTask, err := h.svc.GetTask(req.Id)
	if err != nil {
		return nil, err
	}

	// Проверяем пользователя
	if _, err := h.userClient.GetUser(ctx, &userpb.User{Id: currentTask.UserID}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", currentTask.UserID, err)
	}

	// Обновляем задачу
	t, err := h.svc.UpdateTask(req.Id, req.Title, req.IsDone)
	if err != nil {
		return nil, err
	}

	return &taskpb.Task{
		Id:     t.ID,
		UserId: t.UserID,
		Title:  t.Title,
		IsDone: t.IsDone,
	}, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	err := h.svc.DeleteTask(req.Id)
	if err != nil {
		return nil, err
	}
	return &taskpb.DeleteTaskResponse{Ok: true}, nil
}