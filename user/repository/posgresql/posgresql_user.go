package posgresql

import (
	"context"
	"eventzezz_backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type posgresqlUserRepository struct {
	conn *gorm.DB
}

func NewPosgresqlUserRepository(conn *gorm.DB) domain.UserRepository {
	return &posgresqlUserRepository{conn}
}

func (m posgresqlUserRepository) GetUsers(ctx context.Context, cursor string, num int64) (res []domain.User, nextCursor string, err error) {
	var users []domain.User
	result := m.conn.Find(&users)
	return users, "", result.Error
}

func (m posgresqlUserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var user domain.User
	result := m.conn.Where("id = ?", id).First(&user)
	return user, result.Error

}

func (m posgresqlUserRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	result := m.conn.First(&user, "email = ?", email)
	return user, result.Error
}

func (m posgresqlUserRepository) UpdateUser(ctx context.Context, u *domain.User, id uuid.UUID) error {
	var user domain.User
	result := m.conn.Model(&user).Where("id = ?", id).Updates(u)
	if result.Error != nil {
		panic(result.Error)
	}
	return result.Error
}

func (m posgresqlUserRepository) CreateUser(ctx context.Context, u *domain.User) error {
	result := m.conn.Create(u)
	return result.Error
}

func (m posgresqlUserRepository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	var user domain.User
	result := m.conn.Where("id = ?", id).Delete(&user)
	return result.Error
}
