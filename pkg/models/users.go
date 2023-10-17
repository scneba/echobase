package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `gorm:"column:id"`
	FirstName    string     `gorm:"column:first_name"`
	LastName     string     `gorm:"column:last_name"`
	Username     string     `gorm:"column:user_name"`
	Email        string     `gorm:"column:email"`
	Address      string     `gorm:"column:address"`
	PhoneNumber  string     `gorm:"column:phone_number"`
	Date_Created *time.Time `gorm:"column:date_created_utc"`
	Date_Updated *time.Time `gorm:"column:date_updated_utc"`
}

func (User) TableName() string {
	return "accounts.users"
}
