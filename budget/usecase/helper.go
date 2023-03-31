package usecase

import "eventzezz_backend/domain"

func ToBudgetResponses(budgets []domain.Budget) []domain.BudgetResp {
	var budgetResponses []domain.BudgetResp
	for _, budget := range budgets {
		budgetResponse := ToBudgetResponse(budget)
		budgetResponses = append(budgetResponses, budgetResponse)
	}
	return budgetResponses
}

func ToBudgetResponse(budget domain.Budget) domain.BudgetResp {
	budgetResponse := domain.BudgetResp{
		ID:        budget.ID,
		Purpose:   budget.Purpose,
		Amount:    budget.Amount,
		Note:      budget.Note,
		CreatedAt: budget.CreatedAt,
		UpdatedAt: budget.UpdatedAt,
	}
	return budgetResponse
}
