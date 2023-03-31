package usecase

import (
	"context"
	"eventzezz_backend/domain"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type taskUsecase struct {
	taskRepo       domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(u domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &taskUsecase{
		taskRepo:       u,
		contextTimeout: timeout,
	}
}

func (t taskUsecase) GetTasks(ctx context.Context, cursor string, num int64) ([]domain.TaskResponse, string, error) {
	tasks, nextCursor, err := t.taskRepo.GetTasks(ctx, cursor, num)
	var taskResponses []domain.TaskResponse
	for _, task := range tasks {
		taskCategoryAssignment, _ := t.taskRepo.GetTaskCategoryAssignmentsByTaskID(ctx, task.ID)
		taskResponse := ToTaskResponse(task, taskCategoryAssignment)
		taskResponses = append(taskResponses, taskResponse)
	}
	return taskResponses, nextCursor, err
}

func (t taskUsecase) GetTaskByID(ctx context.Context, id uuid.UUID) (domain.TaskResponse, error) {
	task, err := t.taskRepo.GetTaskByID(ctx, id)
	taskCategoryAssignment, err := t.taskRepo.GetTaskCategoryAssignmentsByTaskID(ctx, id)
	fmt.Println(taskCategoryAssignment)
	taskResponse := ToTaskResponse(task, taskCategoryAssignment)

	return taskResponse, err
}

func (t taskUsecase) UpdateTask(ctx context.Context, u *domain.Task, id uuid.UUID) error {
	err := t.taskRepo.UpdateTask(ctx, u, id)
	return err
}

func (t taskUsecase) CreateTask(ctx context.Context, task *domain.Task) error {

	err := t.taskRepo.CreateTask(ctx, task)
	return err
}

func (t taskUsecase) DeleteTaskByID(ctx context.Context, id uuid.UUID) error {
	err := t.taskRepo.DeleteTaskByID(ctx, id)
	return err
}
