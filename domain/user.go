package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username  string    `json:"username" gorm:"type:varchar(255);uniqueIndex"`
	Email     string    `json:"email" gorm:"type:varchar(255);uniqueIndex"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	Role      string    `json:"role"  gorm:"type:varchar(255)"`
	Token     string    `json:"token" gorm:"-"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserUsecase interface {
	GetUsers(ctx context.Context, cursor string, num int64) ([]User, string, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)

	UpdateUser(ctx context.Context, u *User, id uuid.UUID) error
	CreateUser(context.Context, *User) error
	DeleteUserByID(ctx context.Context, id uuid.UUID) error
}

type UserRepository interface {
	GetUsers(ctx context.Context, cursor string, num int64) (res []User, nextCursor string, err error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)

	UpdateUser(ctx context.Context, u *User, id uuid.UUID) error
	CreateUser(ctx context.Context, u *User) error
	DeleteUserByID(ctx context.Context, id uuid.UUID) error
}
