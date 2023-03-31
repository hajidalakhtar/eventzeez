package seeder

import (
	"eventzezz_backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func CreateEvent(db *gorm.DB,
	ID uuid.UUID,
	eventName string,
	eventDate time.Time,
	eventTime string,
	authorID uuid.UUID,
	location string,
	description string,
	eventType string,
) error {
	return db.Create(&domain.Event{
		ID:          ID,
		EventName:   eventName,
		EventDate:   eventDate,
		EventTime:   eventTime,
		AuthorID:    authorID,
		Location:    location,
		Description: description,
		Type:        eventType,
	}).Error
}

func CreateEventUser(db *gorm.DB, eventID uuid.UUID, userID uuid.UUID) error {
	return db.Create(&domain.EventUserAssignment{
		EventID: eventID,
		UserID:  userID,
	}).Error
}
