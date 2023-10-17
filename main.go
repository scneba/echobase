package main

import (
	"fmt"
	"net/http"
	"strconv"

	"gobase.com/base/config"

	"gobase.com/base/pkg/database"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	godotenv.Load()
	conf := config.BuildConfigs()

	// connect to database
	fmt.Println("Connecting to database")
	if err := database.Connect("postgres", conf.ConnectionString); err != nil {
		panic("Unable to connect database " + conf.ConnectionString)
	}
	defer database.CloseDatabaseConnection()

	// run database migrations
	fmt.Println("Running database migrations...")
	migrations := &migrate.FileMigrationSource{
		Dir: "migrate/migrations",
	}
	migrate.SetTable("go_base_migrations")
	n, err := migrate.Exec(database.Db.DB(), "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(strconv.Itoa(n) + " migrations applied!")

	service := initializeServices(database.Db)

	router := mux.NewRouter()
	initializeRoutes(router, service)

	http.ListenAndServe(":5010", router)

}
