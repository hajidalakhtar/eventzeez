package usecase

import (
	"context"
	"eventzezz_backend/domain"
	"eventzezz_backend/helper"
	"github.com/google/uuid"
	"math"
	"time"
)

type rundownUsecase struct {
	rundownRepo    domain.RundownRepository
	contextTimeout time.Duration
}

func NewRundownUsecase(u domain.RundownRepository, timeout time.Duration) domain.RundownUsecase {
	return &rundownUsecase{
		rundownRepo:    u,
		contextTimeout: timeout,
	}
}

func (ru rundownUsecase) GetRundownByID(ctx context.Context, id uuid.UUID) (domain.RundownResp, error) {
	resp, err := ru.rundownRepo.GetRundownByID(ctx, id)
	result := ToRoundownResponse(resp)
	return result, err
}

func (ru rundownUsecase) GetRundownsByEventId(ctx context.Context, eventID uuid.UUID, page int, limit int) ([]domain.RundownResp, domain.PaginatedResponse, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	resp, totalItems, err := ru.rundownRepo.GetRundownsByEventId(ctx, eventID, page, limit)
	result := ToRundownResponses(resp)

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))
	var nextPage int
	if len(resp) == limit {
		nextPage = page + 1
	}

	paginate := helper.ToPaginatedResponse(totalItems, totalPages, page, nextPage, 0)

	return result, paginate, err
}

func (ru rundownUsecase) CreateRundown(ctx context.Context, r *domain.Rundown) error {
	err := ru.rundownRepo.CreateRundown(ctx, r)
	return err
}

func (ru rundownUsecase) UpdateRundown(ctx context.Context, r *domain.Rundown) error {
	err := ru.rundownRepo.UpdateRundown(ctx, r)
	return err
}

func (ru rundownUsecase) DeleteRundownByID(ctx context.Context, id uuid.UUID) error {
	err := ru.rundownRepo.DeleteRundownByID(ctx, id)
	return err
}
