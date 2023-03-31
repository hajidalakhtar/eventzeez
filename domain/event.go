package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	EventName   string    `json:"event_name" gorm:"type:varchar(255);"`
	EventDate   time.Time `json:"event_date" gorm:"type:date;"`
	EventTime   string    `json:"event_time" `
	Location    string    `json:"location" gorm:"type:varchar(255);"`
	Description string    `json:"description" gorm:"type:varchar(255);"`
	Type        string    `json:"type" gorm:"type:varchar(255);"`
	AuthorID    uuid.UUID `json:"author_id" gorm:"type:uuid;"`
	Author      User      `json:"author" gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type EventUserAssignment struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	EventID   uuid.UUID `json:"event_id" gorm:"type:uuid;"`
	Event     Event     `json:"event" gorm:"constraint:OnDelete:CASCADE"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;"`
	User      User      `json:"user" gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type EventResp struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	EventName   string    `json:"event_name" gorm:"type:varchar(255);"`
	EventDate   time.Time `json:"event_date" gorm:"type:date;"`
	EventTime   string    `json:"event_time" `
	Location    string    `json:"location" gorm:"type:varchar(255);"`
	Description string    `json:"description" gorm:"type:varchar(255);"`
	Type        string    `json:"type" gorm:"type:varchar(255);"`
	AuthorID    uuid.UUID `json:"author_id" gorm:"type:uuid;"`
	Author      User      `json:"author" gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type EventRepository interface {
	GetEvents(ctx context.Context, cursor string, num int64, eventID uuid.UUID) ([]Event, string, error)
	GetEventByID(ctx context.Context, id uuid.UUID) (Event, error)
	UpdateEvent(ctx context.Context, u *Event, id uuid.UUID) error
	CreateEvent(context.Context, *Event) (*Event, error)
	DeleteEventByID(ctx context.Context, id uuid.UUID) error
}

type EventUsecase interface {
	GetEvents(ctx context.Context, cursor string, num int64, eventID uuid.UUID) ([]Event, string, error)
	GetEventByID(ctx context.Context, id uuid.UUID) (Event, error)
	UpdateEvent(ctx context.Context, u *Event, id uuid.UUID) error
	CreateEvent(context.Context, *Event) (*Event, error)
	DeleteEventByID(ctx context.Context, id uuid.UUID) error
}
