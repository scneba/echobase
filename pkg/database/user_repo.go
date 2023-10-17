package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"gobase.com/base/pkg/models"
	"gobase.com/base/pkg/registering"
)

type repo struct{}

func (repo *DBRepo) RegisterUser(user registering.User) (id uuid.UUID, err error) {
	userModel := models.User{ID: newUUID()}
	userModel.Address = user.Address
	userModel.Username = user.Username
	userModel.FirstName = user.FirstName
	userModel.LastName = user.LastName
	userModel.Email = user.Email
	userModel.PhoneNumber = user.PhoneNumber

	err = repo.db.Create(&userModel).Error
	return userModel.ID, err
}

func (repo *DBRepo) UserExists(user registering.User) (unique bool, err error) {
	result := repo.db.Where("user_name= ?", user.Username).Or("phone_number= ?", user.PhoneNumber).Or("email = ?", user.Email).Take(&models.User{})

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}
