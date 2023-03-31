package seeder

import (
	"eventzezz_backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateBudget(db *gorm.DB,
	ID uuid.UUID,
	eventID uuid.UUID,
	Purpose string,
	amount float64,
	note string,
) error {
	return db.Create(&domain.Budget{
		ID:      ID,
		EventID: eventID,
		Purpose: Purpose,
		Amount:  amount,
		Note:    note,
	}).Error
}
