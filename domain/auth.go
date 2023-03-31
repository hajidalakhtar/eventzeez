package domain

import (
	"context"
	"github.com/google/uuid"
)

type AuthUsecase interface {
	Login(ctx context.Context, email string, password string) (User, bool, error)
	GetMe(ctx context.Context, id uuid.UUID) (User, error)
	Register(ctx context.Context, username string, password string, email string) (User, error)
	Logout(ctx context.Context, token string) error
	Validate(ctx context.Context, token string) error
}

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (User, error)
}
