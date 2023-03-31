package seeder

import (
	"eventzezz_backend/domain"
	"eventzezz_backend/helper"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB,
	ID uuid.UUID,
	username string,
	email string,
	role string,
	password string,
) error {
	hashedPassword, _ := helper.PasswordEncrypt(password)
	return db.Create(&domain.User{
		ID:       ID,
		Username: username,
		Role:     role,
		Password: hashedPassword,
		Email:    email,
	}).Error
}
