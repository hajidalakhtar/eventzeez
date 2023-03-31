package posgresql

import (
	"context"
	"eventzezz_backend/domain"
	"gorm.io/gorm"
)

type posgresqlAuthRepository struct {
	conn *gorm.DB
}

func NewPosgresqlAuthRepository(conn *gorm.DB) domain.AuthRepository {
	return &posgresqlAuthRepository{conn}
}

func (ar posgresqlAuthRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	result := ar.conn.Where("email = ?", email).First(&user)
	return user, result.Error
}
