package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Participant struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FullName string    `json:"full_name" gorm:"type:varchar(255);"`
	Email    string    `json:"email" gorm:"type:varchar(255);uniqueIndex"`
	Phone    string    `json:"phone" gorm:"type:varchar(255);uniqueIndex"`
}

type ParticipantEvent struct {
	ID            uuid.UUID  `json:"id"`
	ParticipantID uuid.UUID  `json:"participant_id" gorm:"type:varchar(255);"`
	FullName      string     `json:"full_name" gorm:"type:varchar(255);"`
	Email         string     `json:"email" gorm:"type:varchar(255);uniqueIndex"`
	Phone         string     `json:"phone" gorm:"type:varchar(255);uniqueIndex"`
	AttendanceAt  *time.Time `json:"attendance_at"`
}

type ParticipantEventAssignment struct {
	ID            uuid.UUID   `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ParticipantID uuid.UUID   `json:"participant_id" gorm:"type:uuid;"`
	Participant   Participant `json:"participant"  gorm:"constraint:OnDelete:CASCADE"`
	EventID       uuid.UUID   `json:"event_id" gorm:"type:uuid;"`
	Event         Event       `json:"event"  gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt     time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	AttendanceAt  *time.Time  `json:"attendance_at"`
}

type ParticipantUsecase interface {
	GetParticipantsByEventID(ctx context.Context, eventid uuid.UUID) ([]ParticipantEvent, error)
	GetParticipantByID(ctx context.Context, id uuid.UUID) (Participant, error)
	CreateParticipant(ctx context.Context, p *Participant) error
	DeleteParticipantEventByID(ctx context.Context, participantEventID []uuid.UUID) error

	CheckParticipantRegister(ctx context.Context, email string, eventID uuid.UUID) (Participant, string, error)
	JoinEvent(ctx context.Context, p *Participant, eventID uuid.UUID) error
	ClockInEvent(ctx context.Context, joinEventID uuid.UUID) (Participant, error)
	GetAttendanceByEventID(ctx context.Context, eventid uuid.UUID) ([]ParticipantEvent, error)
}

type ParticipantRepository interface {
	GetParticipantsByEventID(ctx context.Context, eventid uuid.UUID) ([]ParticipantEventAssignment, error)
	GetParticipantByID(ctx context.Context, id uuid.UUID) (Participant, error)
	GetParticipantByEmail(ctx context.Context, email string) (Participant, error)
	CreateParticipant(ctx context.Context, p *Participant) (uuid.UUID, error)
	DeleteParticipantEventByID(ctx context.Context, participantEventID []uuid.UUID) error

	CheckParticipantData(ctx context.Context, p *Participant) (bool, uuid.UUID, error)
	CheckParticipantRegister(ctx context.Context, participantID uuid.UUID, eventID uuid.UUID) (bool, error)
	JoinEvent(ctx context.Context, participantID uuid.UUID, eventID uuid.UUID) error
	GetAttendanceByEventID(ctx context.Context, eventid uuid.UUID) ([]ParticipantEventAssignment, error)
	UpdateAttendanceDateForParticipantEvent(ctx context.Context, joinEventID uuid.UUID) (ParticipantEventAssignment, error)
}
