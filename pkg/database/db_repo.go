package database

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type DBRepo struct {
	db *gorm.DB
}

func NewDBRepo(db *gorm.DB) DBRepo {
	return DBRepo{db}
}

func newUUID() uuid.UUID {
	uuid, _ := uuid.NewUUID()
	return uuid
}
