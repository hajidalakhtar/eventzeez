package seeder

import (
	"eventzezz_backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateParticipant(db *gorm.DB,
	id uuid.UUID,
	fullName string,
	email string,
	phone string,
) error {
	return db.Create(&domain.Participant{
		ID:       id,
		FullName: fullName,
		Email:    email,
		Phone:    phone,
	}).Error
}

func CreateParticipantEvent(db *gorm.DB,
	participantID uuid.UUID,
	eventID uuid.UUID,
) error {
	return db.Create(&domain.ParticipantEventAssignment{
		ParticipantID: participantID,
		EventID:       eventID,
	}).Error
}
