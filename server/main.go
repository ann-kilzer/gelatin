package main

import (
	"database/sql"
	"net/http"

	"github.com/ann-kilzer/gelatin/app"
	"github.com/ann-kilzer/gelatin/database"
	"github.com/ann-kilzer/gelatin/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := database.ConnectToDB(database.DefaultConfig())
	if err != nil {
		panic(err)
	}

	e := initEcho(db)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func initEcho(db *sql.DB) *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Pass DB into custom context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(app.NewContext(db, c))
		}
	})

	// Routes
	e.GET("/", root)

	e.POST("/locations", handlers.CreateLocation)
	e.GET("/locations/:id", handlers.GetLocation)
	e.PUT("/locations/:id", handlers.UpdateLocation)
	e.GET("/locations", handlers.GetLocations)

	return e
}

// Handler
func root(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
