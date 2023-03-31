package helper

import (
	"eventzezz_backend/domain"
	"golang.org/x/crypto/bcrypt"
)

func PasswordEncrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func PasswordCompare(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func ToPaginatedResponse(TotalItems int64, TotalPages int, CurrentPage int, NextPage int, PrevPage int) domain.PaginatedResponse {
	paginate := domain.PaginatedResponse{
		TotalItems:  TotalItems,
		TotalPages:  TotalPages,
		CurrentPage: CurrentPage,
		NextPage:    NextPage,
		PrevPage:    PrevPage,
	}
	return paginate
}
