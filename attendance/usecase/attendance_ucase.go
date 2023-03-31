package usecase

import (
	"context"
	"errors"
	"eventzezz_backend/domain"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type attendanceUsecase struct {
	attendanceRepo  domain.AttendanceRepository
	participantRepo domain.ParticipantRepository
	contextTimeout  time.Duration
}

func NewAttendanceUsecase(u domain.AttendanceRepository, p domain.ParticipantRepository, timeout time.Duration) domain.AttendanceUsecase {
	return &attendanceUsecase{
		attendanceRepo:  u,
		participantRepo: p,
		contextTimeout:  timeout,
	}
}

func (au attendanceUsecase) CheckParticipantRegister(ctx context.Context, email string, eventID uuid.UUID) (domain.Participant, string, error) {
	res, err := au.participantRepo.GetParticipantByEmail(ctx, email)
	if res.Email == "" {
		return res, "participant not found", err

	}

	isExist, err := au.attendanceRepo.CheckParticipantRegister(ctx, res.ID, eventID)
	if isExist {
		return res, "participant already register", err
	} else {
		return res, "participant not register", err

	}

}

func (au attendanceUsecase) JoinEvent(ctx context.Context, p *domain.Participant, eventID uuid.UUID) error {
	isExist, participantID, err := au.attendanceRepo.CheckParticipantData(ctx, p)
	if !isExist {
		participantID, err = au.participantRepo.CreateParticipant(ctx, p)

	}

	fmt.Println(participantID, eventID)
	isRegist, err := au.attendanceRepo.CheckParticipantRegister(ctx, participantID, eventID)
	if isRegist {
		err = errors.New("participant already register")
		return err
	}
	err = au.attendanceRepo.JoinEvent(ctx, participantID, eventID)
	return err
}

func (au attendanceUsecase) ClockInEvent(ctx context.Context, joinEventID uuid.UUID) (domain.Participant, error) {
	res, err := au.attendanceRepo.UpdateAttendanceDateForParticipantEvent(ctx, joinEventID)
	return res.Participant, err
}

func (au attendanceUsecase) GetAttendanceByEventID(ctx context.Context, eventid uuid.UUID) ([]domain.ParticipantEvent, error) {
	res, err := au.attendanceRepo.GetAttendanceByEventID(ctx, eventid)
	participants := EventParticipantToParticipant(res)
	return participants, err
}
