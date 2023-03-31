package posgresql

import (
	"context"
	"eventzezz_backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type posgresqlBudgetRepository struct {
	conn *gorm.DB
}

func NewPosgresqlBudgetRepository(conn *gorm.DB) domain.BudgetRepository {
	return &posgresqlBudgetRepository{conn}
}

func (p posgresqlBudgetRepository) GetBudgetsByEventID(ctx context.Context, cursor string, num int64, id uuid.UUID) ([]domain.Budget, string, error) {
	var budget []domain.Budget
	result := p.conn.Where("event_id = ?", id).Find(&budget)
	return budget, "", result.Error
}

func (p posgresqlBudgetRepository) GetBudgetByID(ctx context.Context, id uuid.UUID) (domain.Budget, error) {
	var budget domain.Budget
	result := p.conn.First(&budget, id)
	return budget, result.Error
}

func (p posgresqlBudgetRepository) UpdateBudget(ctx context.Context, u *domain.Budget, id uuid.UUID) error {
	result := p.conn.Model(&domain.Budget{}).Where("id = ?", id).Updates(u)
	return result.Error
}

func (p posgresqlBudgetRepository) CreateBudget(ctx context.Context, budget *domain.Budget) error {
	result := p.conn.Create(budget)
	return result.Error
}

func (p posgresqlBudgetRepository) DeleteBudgetByID(ctx context.Context, id uuid.UUID) error {
	result := p.conn.Delete(&domain.Budget{}, id)
	return result.Error
}
