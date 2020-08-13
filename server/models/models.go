package models

import (
	"database/sql"

	"github.com/ann-kilzer/gelatin/app"
	"github.com/labstack/echo/v4"
)

func dbFromContext(c echo.Context) *sql.DB {
	return c.(*app.Context).DB
}
