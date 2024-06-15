package main

import (
	"github.com/traP-jp/h24s_10/api"
	"github.com/traP-jp/h24s_10/handler"
	"github.com/traP-jp/h24s_10/migration"
	"github.com/traP-jp/h24s_10/model"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// connect to database
	db, err := sqlx.Connect("mysql", model.MySQL().FormatDSN())
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	// migrate tables
	if err := migration.MigrateTables(db.DB); err != nil {
		e.Logger.Fatal(err)
	}

	// setup repository
	repo := model.New(db)

	// setup routes
	h := handler.New(repo)

	api.RegisterHandlersWithBaseURL(e, h, "/api")

	e.Logger.Fatal(e.Start(":8080"))
}
