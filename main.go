package main

import (
	"fmt"
	"strconv"

	"gobase.com/base/config"

	"gobase.com/base/pkg/database"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	//router := mux.NewRouter()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	initializeRoutes(e, service)
	e.Logger.Fatal(e.Start(":5010"))

}
