package usecase

import (
	"context"
	"errors"
	"eventzezz_backend/domain"
	"eventzezz_backend/user/repository/posgresql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestUserUsecase_GetUserByID(t *testing.T) {
	mockUserRepo := new(posgresql.UserRepositoryMock)
	userID := uuid.New()
	user := domain.User{
		ID:        userID,
		Username:  "user1",
		Email:     "user1@example.com",
		Password:  "password1",
		Role:      "user",
		Token:     "nil",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.Mock.On("GetUserByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(user, nil).Once()
		u := NewUserUsecase(mockUserRepo, time.Second*2)
		_, err := u.GetUserByID(context.TODO(), userID)
		assert.NoError(t, err)

	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.Mock.On("GetUserByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(domain.User{}, errors.New("Unexpected")).Once()

		u := NewUserUsecase(mockUserRepo, time.Second*2)

		a, err := u.GetUserByID(context.TODO(), userID)

		assert.Error(t, err)
		assert.Equal(t, domain.User{}, a)

	})
}

func TestUserUsecase_GetUserByEmail(t *testing.T) {
	mockUserRepo := new(posgresql.UserRepositoryMock)
	userID := uuid.New()
	user := domain.User{
		ID:        userID,
		Username:  "user1",
		Email:     "user1@example.com",
		Password:  "password1",
		Role:      "user",
		Token:     "nil",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.Mock.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(user, nil).Once()
		u := NewUserUsecase(mockUserRepo, time.Second*2)
		_, err := u.GetUserByEmail(context.TODO(), user.Email)
		assert.NoError(t, err)

	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.Mock.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(domain.User{}, errors.New("Unexpected")).Once()

		u := NewUserUsecase(mockUserRepo, time.Second*2)
		_, err := u.GetUserByEmail(context.TODO(), "fail@example.com")
		assert.Error(t, err)
		//assert.Equal(t, domain.User{}, a)

	})
}

func TestUserUsecase_PostCreateUser(t *testing.T) {
	mockUserRepo := new(posgresql.UserRepositoryMock)
	userID := uuid.New()
	user := domain.User{
		ID:        userID,
		Username:  "user1",
		Email:     "user1@example.com",
		Password:  "password1",
		Role:      "user",
		Token:     "nil",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	t.Run("success", func(t *testing.T) {
		tempMockUser := user
		mockUserRepo.Mock.On("CreateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		u := NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.CreateUser(context.TODO(), &tempMockUser)

		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		tempMockUser := user
		mockUserRepo.Mock.On("CreateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(errors.New("Duplicate")).Once()

		u := NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.CreateUser(context.TODO(), &tempMockUser)

		assert.Error(t, err)
	})
}

func TestUserUsecase_UpdateUser(t *testing.T) {
	mockUserRepo := new(posgresql.UserRepositoryMock)
	userID := uuid.New()
	user := domain.User{
		ID:        userID,
		Username:  "user1",
		Email:     "user1@example.com",
		Password:  "password1",
		Role:      "user",
		Token:     "nil",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.Mock.On("UpdateUser", mock.Anything, mock.AnythingOfType("*domain.User"), mock.AnythingOfType("uuid.UUID")).Return(nil).Once()

		u := NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.UpdateUser(context.TODO(), &user, userID)

		assert.NoError(t, err)

	})

	t.Run("fail", func(t *testing.T) {
		mockUserRepo.Mock.On("UpdateUser", mock.Anything, mock.AnythingOfType("*domain.User"), mock.AnythingOfType("uuid.UUID")).Return(errors.New("Unexpected")).Once()

		u := NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.UpdateUser(context.TODO(), &user, userID)

		assert.Error(t, err)
	})
}

func TestUserUsecase_DeleteUser(t *testing.T) {
	mockUserRepo := new(posgresql.UserRepositoryMock)
	userID := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockUserRepo.Mock.On("DeleteUserByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil).Once()
		u := NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.DeleteUserByID(context.TODO(), userID)
		assert.NoError(t, err)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.Mock.On("DeleteUserByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(errors.New("Unexpected")).Once()
		u := NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.DeleteUserByID(context.TODO(), userID)
		assert.Error(t, err)
	})

}

func TestUserUsecase_GetUsers(t *testing.T) {
	mockUserRepo := new(posgresql.UserRepositoryMock)
	userID := uuid.New()
	user1 := domain.User{
		ID:        userID,
		Username:  "user1",
		Email:     "user1@example.com",
		Password:  "password1",
		Role:      "user",
		Token:     "nil",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	user2 := domain.User{
		ID:        uuid.New(),
		Username:  "user2",
		Email:     "user2@example.com",
		Password:  "password2",
		Role:      "admin",
		Token:     "nil",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	users := []domain.User{user1, user2}
	t.Run("success", func(t *testing.T) {
		mockUserRepo.Mock.On("GetUsers", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(users, "next-cursor", nil).Once()

		u := NewUserUsecase(mockUserRepo, time.Second*2)
		num := int64(1)
		cursor := "12"

		list, _, err := u.GetUsers(context.TODO(), cursor, num)
		assert.NoError(t, err)
		assert.Len(t, list, len(users))
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.Mock.On("GetUsers", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(nil, "", errors.New("Unexpexted Error")).Once()

		u := NewUserUsecase(mockUserRepo, time.Second*2)
		num := int64(1)
		cursor := "12"
		_, _, err := u.GetUsers(context.TODO(), cursor, num)

		assert.Error(t, err)
	})
}
