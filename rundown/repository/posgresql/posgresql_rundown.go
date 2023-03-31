package posgresql

import (
	"context"
	"eventzezz_backend/domain"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type posgresqlRundownRepository struct {
	conn *gorm.DB
}

func NewPosgresqlRundownRepository(conn *gorm.DB) domain.RundownRepository {
	return &posgresqlRundownRepository{conn}
}

func (p posgresqlRundownRepository) GetRundownByID(ctx context.Context, id uuid.UUID) (domain.Rundown, error) {
	var rundown domain.Rundown
	result := p.conn.Where("id = ?", id).First(&rundown)
	return rundown, result.Error
}

func (p posgresqlRundownRepository) CreateRundown(ctx context.Context, r *domain.Rundown) error {
	result := p.conn.Create(r)
	return result.Error
}

func (p posgresqlRundownRepository) GetRundownsByEventId(ctx context.Context, eventID uuid.UUID, page int, limit int) ([]domain.Rundown, int64, error) {
	var rundowns []domain.Rundown
	var totalItems int64

	offset := (page - 1) * limit
	_ = p.conn.Model(&rundowns).Where("event_id", eventID).Count(&totalItems).Error
	result := p.conn.Where("event_id", eventID).Limit(limit).Offset(offset).Find(&rundowns)

	return rundowns, totalItems, result.Error
}

func (p posgresqlRundownRepository) UpdateRundown(ctx context.Context, r *domain.Rundown) error {

	result := p.conn.Model(&domain.Rundown{}).Where("id = ?", r.ID).Updates(r)
	fmt.Println(result.Error)
	return result.Error
}

func (p posgresqlRundownRepository) DeleteRundownByID(ctx context.Context, id uuid.UUID) error {
	result := p.conn.Where("id = ?", id).Delete(&domain.Rundown{})
	return result.Error
}
