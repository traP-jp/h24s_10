package main

import (
	"os"
	"strconv"

	"github.com/traP-jp/h24s_10/api"
	"github.com/traP-jp/h24s_10/handler"
	"github.com/traP-jp/h24s_10/migration"
	"github.com/traP-jp/h24s_10/model"
	"github.com/traP-jp/h24s_10/traqclient"
	"github.com/traPtitech/go-traq"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	dev, err := strconv.ParseBool(os.Getenv("DEVELOPMENT"))
	if err != nil {
		dev = false
	}
	if !dev {
		e.Use(handler.TraQIDMiddleware)
	} else {
		e.Use(handler.DevTraQIDMiddleware)
	}

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

	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)

	traqclient := traqclient.New(apiClient)

	// setup routes
	h := handler.New(repo, traqclient)

	api.RegisterHandlersWithBaseURL(e, h, "/api")

	e.Logger.Fatal(e.Start(":8080"))
}
