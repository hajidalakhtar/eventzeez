package usecase

import (
	"context"
	"errors"
	"eventzezz_backend/domain"
	"fmt"
	"github.com/google/uuid"
)

type participantUsecase struct {
	participantRepo domain.ParticipantRepository
}

func NewParticipantUsecase(p domain.ParticipantRepository) domain.ParticipantUsecase {
	return &participantUsecase{
		participantRepo: p,
	}
}

func (pu participantUsecase) GetParticipantsByEventID(ctx context.Context, eventID uuid.UUID) ([]domain.ParticipantEvent, error) {
	res, err := pu.participantRepo.GetParticipantsByEventID(ctx, eventID)
	participans := EventParticipantToParticipant(res)
	return participans, err
}

// TODO: need to refactor this
func (pu participantUsecase) CheckParticipantRegister(ctx context.Context, email string, eventID uuid.UUID) (domain.Participant, string, error) {
	res, err := pu.participantRepo.GetParticipantByEmail(ctx, email)
	if res.Email == "" {
		return res, "participant not found", err

	}

	isExist, err := pu.participantRepo.CheckParticipantRegister(ctx, res.ID, eventID)
	if isExist {
		return res, "participant already register", err
	} else {
		return res, "participant not register", err

	}

}

func (pu participantUsecase) GetParticipantByID(ctx context.Context, id uuid.UUID) (domain.Participant, error) {
	panic("implement me")
}

func (pu participantUsecase) ClockInEvent(ctx context.Context, joinEventID uuid.UUID) (domain.Participant, error) {
	res, err := pu.participantRepo.UpdateAttendanceDateForParticipantEvent(ctx, joinEventID)
	return res.Participant, err
}

func (pu participantUsecase) CreateParticipant(ctx context.Context, p *domain.Participant) error {
	_, err := pu.participantRepo.CreateParticipant(ctx, p)
	return err
}

func (pu participantUsecase) DeleteParticipantEventByID(ctx context.Context, participantEventID []uuid.UUID) error {
	err := pu.participantRepo.DeleteParticipantEventByID(ctx, participantEventID)
	return err
}

func (pu participantUsecase) JoinEvent(ctx context.Context, p *domain.Participant, eventID uuid.UUID) error {
	isExist, participantID, err := pu.participantRepo.CheckParticipantData(ctx, p)
	if !isExist {
		participantID, err = pu.participantRepo.CreateParticipant(ctx, p)

	}

	fmt.Println(participantID, eventID)
	isRegist, err := pu.participantRepo.CheckParticipantRegister(ctx, participantID, eventID)
	if isRegist {
		err = errors.New("participant already register")
		return err
	}
	err = pu.participantRepo.JoinEvent(ctx, participantID, eventID)
	return err
}

func (pu participantUsecase) GetAttendanceByEventID(ctx context.Context, eventID uuid.UUID) ([]domain.ParticipantEvent, error) {

	res, err := pu.participantRepo.GetAttendanceByEventID(ctx, eventID)
	participants := EventParticipantToParticipant(res)
	return participants, err
}
