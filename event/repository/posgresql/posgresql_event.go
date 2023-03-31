package posgresql

import (
	"context"
	"eventzezz_backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type posgresqlEventRepository struct {
	conn *gorm.DB
}

func NewPosgresqlEventRepository(conn *gorm.DB) domain.EventRepository {
	return &posgresqlEventRepository{conn}
}

func (er posgresqlEventRepository) GetEvents(ctx context.Context, cursor string, num int64, userID uuid.UUID) ([]domain.Event, string, error) {
	var events []domain.Event
	result := er.conn.Preload("Author").Where("author_id", userID).Find(&events)
	return events, "", result.Error
}

func (er posgresqlEventRepository) GetEventByID(ctx context.Context, id uuid.UUID) (domain.Event, error) {
	var event domain.Event
	result := er.conn.Where("id = ?", id).First(&event)
	return event, result.Error
}

func (er posgresqlEventRepository) UpdateEvent(ctx context.Context, u *domain.Event, id uuid.UUID) error {
	var event domain.Event
	result := er.conn.Model(&event).Where("id = ?", id).Updates(u)
	return result.Error
}

func (er posgresqlEventRepository) CreateEvent(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	result := er.conn.Create(event)
	return event, result.Error
}

func (er posgresqlEventRepository) DeleteEventByID(ctx context.Context, id uuid.UUID) error {
	var event domain.Event
	result := er.conn.Where("id = ?", id).Delete(&event)
	return result.Error
}
