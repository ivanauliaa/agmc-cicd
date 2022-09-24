package main

import (
	"os"

	"day9-cicd/database"
	"day9-cicd/database/migration"
	"day9-cicd/internal/factory"
	"day9-cicd/internal/http"
	"day9-cicd/internal/middleware"
	"day9-cicd/mongo"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}
}

func main() {
	database.CreateConnection()
	mongo.CreateConnection()

	migration.Migrate()

	f := factory.NewFactory()
	e := echo.New()
	middleware.Init(e)
	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
