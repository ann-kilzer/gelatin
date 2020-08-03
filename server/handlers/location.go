package handlers

import (
	"context"
	"net/http"

	"github.com/ann-kilzer/gelatin/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Location is the json request body for a location you can sleep
type Location struct {
	ID       string `json:"id" form:"name" query:"name"`
	Name     string `json:"name" form:"name" query:"name"`
	Softness int    `json:"softness" form:"email" query:"email"`
}

// CreateLocation creates a new location
func CreateLocation(c echo.Context) error {
	payload := new(Location)
	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var p models.Location
	p.LocationName = payload.Name
	p.Name = payload.LocationName
	p.Softness = payload.Softness
	err := p.Insert(context.TODO(), dbFromContext(c), boil.Infer())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, p)
}

// GetLocation retrieves one location by ID
func GetLocation(c echo.Context) error {
	params, err := readParams(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	location, err := models.FindLocation(context.TODO(), dbFromContext(c), params.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, location)
}

// GetLocations retrieves all locations in the system
func GetLocations(c echo.Context) error {
	locations, err := models.Locations().All(context.TODO(), dbFromContext(c))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, locations)
}

// UpdateLocation updates a location
func UpdateLocation(c echo.Context) error {
	db := dbFromContext(c)
	payload := new(Location)
	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	params, err := readParams(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	p, err := models.FindLocation(context.TODO(), db, params.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	p.Name = payload.LocationName
	p.Softness = payload.Softness
	_, err = p.Update(context.TODO(), db, boil.Infer())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
