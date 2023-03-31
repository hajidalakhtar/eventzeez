package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Rundown struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title       string    `json:"title" gorm:"type:varchar(255);"`
	Description string    `json:"description" gorm:"type:text;"`
	EventID     uuid.UUID `json:"event_id" gorm:"type:uuid;"`
	Event       Event     `json:"event" gorm:"constraint:OnDelete:CASCADE"`
	StartDate   time.Time `json:"start_date" gorm:"type:timestamp;"`
	EndDate     time.Time `json:"end_date" gorm:"type:timestamp;"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp;"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamp;"`
}

type RundownResp struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EventID     uuid.UUID `json:"event_id"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RundownUsecase interface {
	GetRundownByID(ctx context.Context, id uuid.UUID) (RundownResp, error)
	GetRundownsByEventId(ctx context.Context, eventID uuid.UUID, page int, limit int) ([]RundownResp, PaginatedResponse, error)
	CreateRundown(ctx context.Context, r *Rundown) error
	UpdateRundown(ctx context.Context, r *Rundown) error
	DeleteRundownByID(ctx context.Context, id uuid.UUID) error
}

type RundownRepository interface {
	GetRundownByID(ctx context.Context, id uuid.UUID) (Rundown, error)
	GetRundownsByEventId(ctx context.Context, eventID uuid.UUID, page int, limit int) ([]Rundown, int64, error)
	CreateRundown(ctx context.Context, r *Rundown) error
	UpdateRundown(ctx context.Context, r *Rundown) error
	DeleteRundownByID(ctx context.Context, id uuid.UUID) error
}
