package posgresql

import (
	"context"
	"eventzezz_backend/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func NewUserRepositoryMock(mock mock.Mock) domain.UserRepository {
	return &UserRepositoryMock{Mock: mock}
}

func (_m UserRepositoryMock) GetUsers(ctx context.Context, cursor string, num int64) (res []domain.User, nextCursor string, err error) {
	ret := _m.Mock.Called(ctx, cursor, num)

	var r0 []domain.User
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) []domain.User); ok {
		r0 = rf(ctx, cursor, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) string); ok {
		r1 = rf(ctx, cursor, num)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, int64) error); ok {
		r2 = rf(ctx, cursor, num)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2

}

func (_m UserRepositoryMock) GetUserByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	ret := _m.Mock.Called(ctx, id)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) domain.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m UserRepositoryMock) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	ret := _m.Mock.Called(ctx, email)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m UserRepositoryMock) UpdateUser(ctx context.Context, u *domain.User, id uuid.UUID) error {
	ret := _m.Mock.Called(ctx, u, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User, uuid.UUID) error); ok {
		r0 = rf(ctx, u, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m UserRepositoryMock) CreateUser(ctx context.Context, u *domain.User) error {
	ret := _m.Mock.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m UserRepositoryMock) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	ret := _m.Mock.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0

}

//func (repository *UserRepositoryMock) GetUsers(ctx context.Context, cursor string, num int64) (res []domain.User, nextCursor string, err error) {
//	arguments := repository.Mock.Called(ctx, cursor, num)
//	if arguments.Get(0) == nil {
//		return []domain.User{}, "", errors.New("User Not found")
//
//	} else {
//		users := arguments.Get(0).([]domain.User)
//		return users, "", nil
//	}
//}
//
//func (repository *UserRepositoryMock) GetUserByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
//	arguments := repository.Mock.Called(id)
//	if arguments.Get(0) == nil {
//		return domain.User{}, errors.New("User Not found")
//
//	} else {
//		user := arguments.Get(0).(domain.User)
//		return user, nil
//	}
//}
//
//func (repository *UserRepositoryMock) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
//	arguments := repository.Mock.Called(email)
//	if arguments.Get(0) == nil {
//		return domain.User{}, errors.New("User Not found")
//
//	} else {
//		user := arguments.Get(0).(domain.User)
//		return user, nil
//	}
//	//if arguments.Get(0) == nil {
//	//	return domain.User{}, errors.New("User Not found")
//	//
//	//} else {
//	//	user := arguments.Get(0).(domain.User)
//	//	return user, nil
//	//}
//
//}
//
//func (repository *UserRepositoryMock) UpdateUser(ctx context.Context, u *domain.User, id uuid.UUID) error {
//	arguments := repository.Mock.Called(u, id)
//	if arguments.Get(0) == nil {
//		return nil
//	} else {
//		return arguments.Get(0).(error)
//	}
//}
//
//func (repository *UserRepositoryMock) CreateUser(ctx context.Context, u *domain.User) error {
//	arguments := repository.Mock.Called(u)
//	if arguments.Get(0) == nil {
//		return nil
//	} else {
//		return arguments.Get(0).(error)
//	}
//}
//
//func (repository *UserRepositoryMock) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
//	return nil
//
//}
//
////
