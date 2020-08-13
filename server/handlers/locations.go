package handlers

import (
	"context"
	"net/http"

	gm "github.com/ann-kilzer/gelatin/genmodels"
	"github.com/ann-kilzer/gelatin/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Location is the json request body for a location you can sleep
type Location struct {
	ID       int    `json:"id" form:"name" query:"name"`
	Name     string `json:"name" form:"name" query:"name"`
	Softness int    `json:"softness" form:"email" query:"email"`
}

// CreateLocation creates a new location
func CreateLocation(c echo.Context) error {
	payload := new(Location)
	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	p, err := models.CreateLocation(dbFromContext(c), payload.Name, payload.Softness)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, Location{ID: p.ID, Name: p.LocationName, Softness: p.Softness})
}

// GetLocation retrieves one location by ID
func GetLocation(c echo.Context) error {
	params, err := readParams(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	location, err := gm.FindLocation(context.TODO(), dbFromContext(c), params.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, location)
}

// GetLocations retrieves all locations in the system
func GetLocations(c echo.Context) error {
	rawLocations, err := gm.Locations().All(context.TODO(), dbFromContext(c))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	locations := []Location{}
	for _, rl := range rawLocations {
		locations = append(locations, Location{ID: rl.ID, Name: rl.LocationName, Softness: rl.Softness})
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

	p, err := gm.FindLocation(context.TODO(), db, params.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	p.LocationName = payload.Name
	p.Softness = payload.Softness
	_, err = p.Update(context.TODO(), db, boil.Infer())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
