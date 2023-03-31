package usecase

import (
	"context"
	"eventzezz_backend/domain"
	"github.com/google/uuid"
	"time"
)

type eventUsecase struct {
	eventRepo      domain.EventRepository
	contextTimeout time.Duration
}

func NewEventUsecase(u domain.EventRepository, timeout time.Duration) domain.EventUsecase {
	return &eventUsecase{
		eventRepo:      u,
		contextTimeout: timeout,
	}
}

func (e eventUsecase) GetEvents(ctx context.Context, cursor string, num int64, userID uuid.UUID) ([]domain.Event, string, error) {
	response, nextCursor, err := e.eventRepo.GetEvents(ctx, cursor, num, userID)
	return response, nextCursor, err
}

func (e eventUsecase) GetEventByID(ctx context.Context, id uuid.UUID) (domain.Event, error) {
	response, err := e.eventRepo.GetEventByID(ctx, id)
	return response, err
}

func (e eventUsecase) UpdateEvent(ctx context.Context, u *domain.Event, id uuid.UUID) error {
	err := e.eventRepo.UpdateEvent(ctx, u, id)
	return err
}

func (e eventUsecase) CreateEvent(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	event, err := e.eventRepo.CreateEvent(ctx, event)
	return event, err
}

func (e eventUsecase) DeleteEventByID(ctx context.Context, id uuid.UUID) error {
	err := e.eventRepo.DeleteEventByID(ctx, id)
	return err
}
