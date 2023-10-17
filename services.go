package main

import (
	"github.com/jinzhu/gorm"
	"gobase.com/base/pkg/database"
	"gobase.com/base/pkg/registering"
)

type services struct {
	registering registering.RegisteringInterface
}

func initializeServices(db *gorm.DB) *services {
	dbRepo := database.NewDBRepo(db)
	svs := services{}
	svs.registering = registering.NewRegisteringService(&dbRepo)
	return &svs
}
