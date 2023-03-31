package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Budget struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	EventID   uuid.UUID `json:"event_id" gorm:"type:uuid;"`
	Purpose   string    `json:"purpose" gorm:"type:varchar(255);"`
	Event     Event     `json:"event" gorm:"constraint:OnDelete:CASCADE"`
	Amount    float64   `json:"amount" gorm:"type:float;"`
	Note      string    `json:"note" gorm:"type:varchar(255);"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type BudgetResp struct {
	ID        uuid.UUID `json:"id" `
	Purpose   string    `json:"purpose" `
	Amount    float64   `json:"amount" `
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at" `
}

type BudgetRepository interface {
	GetBudgetsByEventID(ctx context.Context, cursor string, num int64, id uuid.UUID) ([]Budget, string, error)
	GetBudgetByID(ctx context.Context, id uuid.UUID) (Budget, error)
	UpdateBudget(ctx context.Context, u *Budget, id uuid.UUID) error
	CreateBudget(context.Context, *Budget) error
	DeleteBudgetByID(ctx context.Context, id uuid.UUID) error
}

type BudgetUsecase interface {
	GetBudgetsByEventID(ctx context.Context, cursor string, num int64, id uuid.UUID) ([]BudgetResp, string, error)
	GetBudgetByID(ctx context.Context, id uuid.UUID) (BudgetResp, error)
	UpdateBudget(ctx context.Context, u *Budget, id uuid.UUID) error
	CreateBudget(context.Context, *Budget) error
	DeleteBudgetByID(ctx context.Context, id uuid.UUID) error
}
