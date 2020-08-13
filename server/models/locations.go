package models

import (
	"context"
	"database/sql"

	gm "github.com/ann-kilzer/gelatin/genmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// CreateLocation makes a new location and inserts it into the database
func CreateLocation(db *sql.DB, name string, softness int) (gm.Location, error) {
	var p gm.Location
	p.LocationName = name
	p.Softness = softness
	return p, p.Insert(context.TODO(), db, boil.Infer())
}
