package usecase

import (
	"context"
	"eventzezz_backend/domain"
	"github.com/google/uuid"
	"time"
)

type budgetUsecase struct {
	budgetRepo     domain.BudgetRepository
	contextTimeout time.Duration
}

func NewBudgetUsecase(u domain.BudgetRepository, timeout time.Duration) domain.BudgetUsecase {
	return &budgetUsecase{
		budgetRepo:     u,
		contextTimeout: timeout,
	}
}

func (b budgetUsecase) GetBudgetsByEventID(ctx context.Context, cursor string, num int64, id uuid.UUID) ([]domain.BudgetResp, string, error) {
	response, _, err := b.budgetRepo.GetBudgetsByEventID(ctx, cursor, num, id)
	budgetResponses := ToBudgetResponses(response)
	return budgetResponses, "", err
}

func (b budgetUsecase) GetBudgetByID(ctx context.Context, id uuid.UUID) (domain.BudgetResp, error) {
	response, err := b.budgetRepo.GetBudgetByID(ctx, id)
	budgetResponse := ToBudgetResponse(response)
	return budgetResponse, err
}

func (b budgetUsecase) UpdateBudget(ctx context.Context, u *domain.Budget, id uuid.UUID) error {
	err := b.budgetRepo.UpdateBudget(ctx, u, id)
	return err
}

func (b budgetUsecase) CreateBudget(ctx context.Context, budget *domain.Budget) error {
	err := b.budgetRepo.CreateBudget(ctx, budget)
	return err
}

func (b budgetUsecase) DeleteBudgetByID(ctx context.Context, id uuid.UUID) error {
	err := b.budgetRepo.DeleteBudgetByID(ctx, id)
	return err
}
