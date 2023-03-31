package usecase

import (
	"eventzezz_backend/domain"
	"fmt"
)

func EventParticipantToParticipant(pe []domain.ParticipantEventAssignment) []domain.ParticipantEvent {
	fmt.Println(pe)
	var participants []domain.ParticipantEvent
	for _, p := range pe {
		participantEvent := domain.ParticipantEvent{
			ID:            p.ID,
			ParticipantID: p.ParticipantID,
			FullName:      p.Participant.FullName,
			Email:         p.Participant.Email,
			Phone:         p.Participant.Phone,
			AttendanceAt:  p.AttendanceAt,
		}
		participants = append(participants, participantEvent)
	}
	return participants
}
