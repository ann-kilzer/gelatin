package handlers

import (
	"database/sql"
	"strconv"

	"github.com/ann-kilzer/gelatin/app"
	"github.com/labstack/echo/v4"
)

// TODO: refactor into model package
func dbFromContext(c echo.Context) *sql.DB {
	return c.(*app.Context).DB
}

type params struct {
	ID int
}

// readParams handles URL params
func readParams(c echo.Context) (*params, error) {
	id, err := readIntParam(c, "id")
	if err != nil {
		return nil, err
	}
	p := &params{
		ID: id,
	}
	return p, nil
}

func readIntParam(c echo.Context, name string) (int, error) {
	rawID := c.Param(name)
	if rawID == "" {
		return 0, nil
	}
	id, err := strconv.Atoi(rawID)
	if err != nil {
		return 0, err
	}
	return id, nil
}
