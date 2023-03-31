package posgresql

import (
	"context"
	"eventzezz_backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type posgresqlAttendanceRepository struct {
	conn *gorm.DB
}

func NewPosgresqlAttendanceRepository(conn *gorm.DB) domain.AttendanceRepository {
	return &posgresqlAttendanceRepository{conn}
}

func (p2 posgresqlAttendanceRepository) CheckParticipantRegister(ctx context.Context, participantID uuid.UUID, eventID uuid.UUID) (bool, error) {
	var participantEvent domain.ParticipantEventAssignment
	result := p2.conn.Where("participant_id = ? AND event_id = ?", participantID, eventID).First(&participantEvent)
	isExist := participantEvent.ID != uuid.Nil
	return isExist, result.Error
}

// TODO: need fix this
func (p2 posgresqlAttendanceRepository) CheckParticipantData(ctx context.Context, p *domain.Participant) (bool, uuid.UUID, error) {
	var participant domain.Participant
	result := p2.conn.Where("email = ? ", p.Email).First(&participant)
	isExist := result.RowsAffected > 0
	return isExist, participant.ID, result.Error

}
func (p2 posgresqlAttendanceRepository) JoinEvent(ctx context.Context, participantID uuid.UUID, eventID uuid.UUID) error {
	result := p2.conn.Create(&domain.ParticipantEventAssignment{EventID: eventID, ParticipantID: participantID})
	return result.Error
}

func (p2 posgresqlAttendanceRepository) GetAttendanceByEventID(ctx context.Context, eventid uuid.UUID) ([]domain.ParticipantEventAssignment, error) {
	var participantEventAssignment []domain.ParticipantEventAssignment
	result := p2.conn.Joins("Participant").Where("event_id = ? AND attendance_at IS NOT NULL", eventid).Find(&participantEventAssignment)
	return participantEventAssignment, result.Error
}
func (p2 posgresqlAttendanceRepository) UpdateAttendanceDateForParticipantEvent(ctx context.Context, joinEventID uuid.UUID) (domain.ParticipantEventAssignment, error) {
	var participantEvent domain.ParticipantEventAssignment
	result := p2.conn.Model(&domain.ParticipantEventAssignment{}).Preload("Participant").Where("id = ?", joinEventID).Update("attendance_at", gorm.Expr("NOW()")).First(&participantEvent)
	return participantEvent, result.Error
}
