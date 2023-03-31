package usecase

import (
	"eventzezz_backend/domain"
)

func ToRundownResponses(rundowns []domain.Rundown) []domain.RundownResp {

	var rundownResp []domain.RundownResp
	for _, rundown := range rundowns {
		result := ToRoundownResponse(rundown)
		rundownResp = append(rundownResp, result)
	}

	return rundownResp

}

func ToRoundownResponse(rundown domain.Rundown) domain.RundownResp {
	rundownResp := domain.RundownResp{
		ID:          rundown.ID,
		Title:       rundown.Title,
		Description: rundown.Description,
		EventID:     rundown.EventID,
		StartDate:   rundown.StartDate,
		EndDate:     rundown.EndDate,
		CreatedAt:   rundown.CreatedAt,
		UpdatedAt:   rundown.UpdatedAt,
	}

	return rundownResp
}
