package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres driver
)

// Config defines variables for the database connection
type Config struct {
	user     string
	password string
	dbname   string
	sslmode  string
}

// DefaultConfig builds a development config
func DefaultConfig() Config {
	return Config{
		user:     "otto",
		password: "letmeinnow",
		dbname:   "gelatin",
		sslmode:  "disable",
	}
}

// TestConfig builds an API testing config
func TestConfig() Config {
	return Config{
		user:     "otto",
		password: "letmeinnow",
		dbname:   "gelatin_test",
		sslmode:  "disable",
	}
}

// ConnectToDB creates the database connection
func ConnectToDB(config Config) (*sql.DB, error) {
	dbStr := fmt.Sprintf("user=%v password=%s dbname=%s sslmode=%s", config.user, config.password, config.dbname, config.sslmode)
	return sql.Open("postgres", dbStr)
}
