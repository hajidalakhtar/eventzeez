package posgresql

import (
	"context"
	"eventzezz_backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type posgresqlParticipantRepository struct {
	conn *gorm.DB
}

func NewPosgresqlParticipantRepository(conn *gorm.DB) domain.ParticipantRepository {
	return &posgresqlParticipantRepository{conn}
}

func (pr posgresqlParticipantRepository) GetParticipantsByEventID(ctx context.Context, eventID uuid.UUID) ([]domain.ParticipantEventAssignment, error) {
	var participantEvent []domain.ParticipantEventAssignment
	result := pr.conn.Joins("Event").Joins("Participant").Where("event_id = ?", eventID).Find(&participantEvent)

	return participantEvent, result.Error
}

func (pr posgresqlParticipantRepository) GetParticipantByEmail(ctx context.Context, email string) (domain.Participant, error) {
	var participant domain.Participant
	result := pr.conn.Where("email = ?", email).First(&participant)
	return participant, result.Error
}

func (pr posgresqlParticipantRepository) GetParticipantByID(ctx context.Context, id uuid.UUID) (domain.Participant, error) {
	//TODO implement me
	panic("implement me")
}

func (pr posgresqlParticipantRepository) CreateParticipant(ctx context.Context, p *domain.Participant) (uuid.UUID, error) {
	result := pr.conn.Create(p)
	return p.ID, result.Error
}

func (pr posgresqlParticipantRepository) DeleteParticipantEventByID(ctx context.Context, participantEvent []uuid.UUID) error {
	result := pr.conn.Where("id IN ?", participantEvent).Delete(&domain.ParticipantEventAssignment{})
	return result.Error
}

func (pr posgresqlParticipantRepository) UpdateAttendanceDateForParticipantEvent(ctx context.Context, joinEventID uuid.UUID) (domain.ParticipantEventAssignment, error) {
	var participantEvent domain.ParticipantEventAssignment
	result := pr.conn.Model(&domain.ParticipantEventAssignment{}).Preload("Participant").Where("id = ?", joinEventID).Update("attendance_at", gorm.Expr("NOW()")).First(&participantEvent)
	return participantEvent, result.Error
}

func (pr posgresqlParticipantRepository) CheckParticipantRegister(ctx context.Context, participantID uuid.UUID, eventID uuid.UUID) (bool, error) {
	var participantEvent domain.ParticipantEventAssignment
	result := pr.conn.Where("participant_id = ? AND event_id = ?", participantID, eventID).First(&participantEvent)
	//fmt.Println(participantID, eventID)
	//fmt.Println(participantEvent.ID != uuid.Nil)
	isExist := participantEvent.ID != uuid.Nil
	return isExist, result.Error
}

// TODO: need fix this
func (pr posgresqlParticipantRepository) CheckParticipantData(ctx context.Context, p *domain.Participant) (bool, uuid.UUID, error) {
	var participant domain.Participant
	result := pr.conn.Where("email = ? ", p.Email).First(&participant)
	isExist := result.RowsAffected > 0
	return isExist, participant.ID, result.Error

}

func (pr posgresqlParticipantRepository) JoinEvent(ctx context.Context, participantID uuid.UUID, eventID uuid.UUID) error {
	result := pr.conn.Create(&domain.ParticipantEventAssignment{EventID: eventID, ParticipantID: participantID})
	return result.Error
}

func (pr posgresqlParticipantRepository) GetAttendanceByEventID(ctx context.Context, eventid uuid.UUID) ([]domain.ParticipantEventAssignment, error) {
	var participantEventAssignment []domain.ParticipantEventAssignment
	result := pr.conn.Joins("Participant").Where("event_id = ? AND attendance_at IS NOT NULL", eventid).Find(&participantEventAssignment)
	return participantEventAssignment, result.Error
}
