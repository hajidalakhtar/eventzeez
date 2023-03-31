package usecase

import (
	"context"
	"eventzezz_backend/domain"
	"eventzezz_backend/helper"
	"github.com/google/uuid"
	"time"
)

type authUsecase struct {
	authRepo       domain.AuthRepository
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewAuthUsecase(u domain.AuthRepository, ur domain.UserRepository, timeout time.Duration) domain.AuthUsecase {
	return &authUsecase{
		authRepo:       u,
		userRepo:       ur,
		contextTimeout: timeout,
	}
}

func (a authUsecase) Login(ctx context.Context, email string, password string) (domain.User, bool, error) {
	user, err := a.userRepo.GetUserByEmail(ctx, email)
	if err == nil {
		isSuccess := helper.PasswordCompare(password, user.Password)
		user.Token, _ = generateToken(user.ID)
		return user, isSuccess, err
	} else {
		return user, false, err
	}
}

func (a authUsecase) Register(ctx context.Context, username string, password string, email string) (domain.User, error) {
	hashedPassword, _ := helper.PasswordEncrypt(password)
	user := domain.User{Username: username, Email: email, Password: hashedPassword, Role: "user"}
	user.Token, _ = generateToken(user.ID)
	err := a.userRepo.CreateUser(ctx, &user)
	return user, err

}

func (a authUsecase) GetMe(ctx context.Context, id uuid.UUID) (domain.User, error) {
	user, err := a.userRepo.GetUserByID(ctx, id)
	return user, err

}

func (a authUsecase) Logout(ctx context.Context, token string) error {
	//TODO implement me
	panic("implement me")
}

func (a authUsecase) Validate(ctx context.Context, token string) error {
	//TODO implement me
	panic("implement me")
}
