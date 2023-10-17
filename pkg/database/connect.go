package database

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

func Connect(dialect, connectionString string) (err error) {
	db, err := gorm.Open(dialect, connectionString)
	if err != nil {
		log.Fatal(err, nil)
	}
	db.DB().SetMaxOpenConns(60)
	db.LogMode(true)
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC() //set any database time to gmt
	}

	Db = db
	return err
}
func CloseDatabaseConnection() {
	Db.Close()
}
