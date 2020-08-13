// This file contains test setup / teardown, testing state like the db and echo server,
// and some structs used in testing

package handlers_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/ann-kilzer/gelatin/app"
	"github.com/ann-kilzer/gelatin/database"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/queries"
)

// e is our echo server with the DB wired in!
var e *echo.Echo
var db *sql.DB

// TestMain allows us to wire up test setup and shutdown
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

// setup runs before all tests in the package
func setup() {
	var err error
	db, err = database.ConnectToDB(database.TestConfig())
	if err != nil {
		panic(err)
	}

	e = echo.New()

	// Pass DB into custom context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(app.NewContext(db, c))
		}
	})
}

// shutdown runs after all tests in the package
func shutdown() {
	db.Close()
}

// cleanupHelper resets the DB to the original state, until we get proper test transactions working
func cleanupHelper(t *testing.T) {
	_, err := queries.Raw("TRUNCATE TABLE locations").Exec(db)
	if err != nil {
		t.Error(err)
	}
}

// responseBody is the JSON struct for a server response
type responseBody struct {
	Message string `json:"message"`
}
