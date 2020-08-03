package app

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

// Context extends echo.Context with the db handle
type Context struct {
	echo.Context
	DB *sql.DB
}

// NewContext builds a new context
func NewContext(db *sql.DB, ec echo.Context) *Context {
	return &Context{
		DB:      db,
		Context: ec,
	}
}
