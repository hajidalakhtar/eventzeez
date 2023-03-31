package seeder

import (
	"eventzezz_backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func CreateRundown(db *gorm.DB,
	id uuid.UUID,
	title string,
	description string,
	eventID uuid.UUID,
	startDate time.Time,
	endDate time.Time,
) error {
	return db.Create(&domain.Rundown{
		ID:          id,
		Title:       title,
		Description: description,
		EventID:     eventID,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}).Error
}
