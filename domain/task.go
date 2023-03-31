package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type TaskResponse struct {
	ID           uuid.UUID      `json:"id"`
	TaskName     string         `json:"task_name"`
	IsDone       bool           `json:"is_done"`
	Description  string         `json:"description"`
	StartDate    time.Time      `json:"start_date"`
	EndDate      time.Time      `json:"end_date"`
	TaskCategory []TaskCategory `json:"task_category"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type Task struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	EventID     uuid.UUID `json:"event_id" gorm:"type:uuid;"`
	Event       Event     `json:"event" gorm:"constraint:OnDelete:CASCADE"`
	TaskName    string    `json:"task_name" gorm:"type:varchar(255);"`
	IsDone      bool      `json:"is_done" gorm:"type:bool"`
	Description string    `json:"description" `
	StartDate   time.Time `json:"start_date" `
	EndDate     time.Time `json:"end_date" `
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type TaskCategoryAssignment struct {
	ID             uuid.UUID    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TaskID         uuid.UUID    `json:"task_id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Task           Task         `json:"task" gorm:"constraint:OnDelete:CASCADE"`
	TaskCategoryID uuid.UUID    `json:"task_category_id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TaskCategory   TaskCategory `json:"task_category" gorm:"constraint:OnDelete:CASCADE"`
}

type TaskCategory struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name  string    `json:"name" gorm:"type:varchar(255);uniqueIndex"`
	Color string    `json:"color" gorm:"type:varchar(255);uniqueIndex"`
}

type TaskRepository interface {
	GetTasks(ctx context.Context, cursor string, num int64) ([]Task, string, error)
	GetTaskByID(ctx context.Context, id uuid.UUID) (Task, error)
	GetTaskCategoryAssignmentsByTaskID(ctx context.Context, id uuid.UUID) ([]TaskCategoryAssignment, error)

	UpdateTask(ctx context.Context, u *Task, id uuid.UUID) error
	CreateTask(context.Context, *Task) error
	DeleteTaskByID(ctx context.Context, id uuid.UUID) error
}

type TaskUsecase interface {
	GetTasks(ctx context.Context, cursor string, num int64) ([]TaskResponse, string, error)
	GetTaskByID(ctx context.Context, id uuid.UUID) (TaskResponse, error)

	UpdateTask(ctx context.Context, u *Task, id uuid.UUID) error
	CreateTask(context.Context, *Task) error
	DeleteTaskByID(ctx context.Context, id uuid.UUID) error
}
