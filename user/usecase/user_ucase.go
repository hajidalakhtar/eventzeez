package usecase

import (
	"context"
	"eventzezz_backend/domain"
	"github.com/google/uuid"
	"time"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(u domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (userUcase userUsecase) GetUsers(ctx context.Context, cursor string, num int64) (res []domain.User, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}
	ctx, cancel := context.WithTimeout(ctx, userUcase.contextTimeout)
	defer cancel()

	res, nextCursor, err = userUcase.userRepo.GetUsers(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	if err != nil {
		nextCursor = ""
	}
	return
}

func (userUcase userUsecase) GetUserByID(ctx context.Context, id uuid.UUID) (res domain.User, err error) {
	res, err = userUcase.userRepo.GetUserByID(ctx, id)
	return
}

func (userUcase userUsecase) GetUserByEmail(ctx context.Context, email string) (res domain.User, err error) {
	res, err = userUcase.userRepo.GetUserByEmail(ctx, email)
	return
}

func (userUcase userUsecase) UpdateUser(ctx context.Context, u *domain.User, id uuid.UUID) (err error) {
	err = userUcase.userRepo.UpdateUser(ctx, u, id)
	return
}

func (userUcase userUsecase) CreateUser(ctx context.Context, user *domain.User) (err error) {
	err = userUcase.userRepo.CreateUser(ctx, user)
	return
}

func (userUcase userUsecase) DeleteUserByID(ctx context.Context, id uuid.UUID) (err error) {
	err = userUcase.userRepo.DeleteUserByID(ctx, id)
	return
}
