package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ann-kilzer/gelatin/app"
	"github.com/ann-kilzer/gelatin/handlers"
	"github.com/ann-kilzer/gelatin/models"
	"github.com/stretchr/testify/assert"
)

// TestGetLocations tests the GET /locations endpoint
func TestGetLocations(t *testing.T) {
	cases := []struct {
		testName  string
		locations []string
		expCode   int
	}{
		{
			testName:  "Server empty",
			locations: []string{},
			expCode:   http.StatusOK,
		},
		{
			testName:  "Multiple locations",
			locations: []string{"cot", "hammock", "bed"},
			expCode:   http.StatusOK,
		},
		{
			testName:  "Locations with spaces in names",
			locations: []string{"papasan chair", "chaise lounge", "sleeping bag"},
			expCode:   http.StatusOK,
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			tGetLocationsHelper(t, tc.locations, tc.expCode)
			cleanupHelper(t)
		})
	}
}

func tGetLocationsHelper(t *testing.T, locations []string, expCode int) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := app.NewContext(db, e.NewContext(req, rec))
	c.SetPath("/locations")

	insertLocationsHelper(t, locations)

	if assert.NoError(t, handlers.GetLocations(c)) {
		assert.Equal(t, expCode, rec.Code)
		actual := getLocationsHelper(t, rec)
		assert.Equal(t, len(locations), len(actual))
		// Ensure all items in actual are in locations
		for _, a := range actual {
			assert.Contains(t, locations, a.Name)
		}
	}
}

// insertLocationsHelper creates valid locations with the following names in the DB.
func insertLocationsHelper(t *testing.T, locations []string) {
	for _, name := range locations {
		_, err := models.CreateLocation(db, name, 0)
		if err != nil {
			t.Errorf("Error setting up test: %v", err)
		}
	}
}

// getLocationsHelper extracts and unmarshals the JSON from the ResponseRecorder, and returns
// the list of locations
func getLocationsHelper(t *testing.T, rec *httptest.ResponseRecorder) []handlers.Location {
	var locations []handlers.Location
	err := json.Unmarshal(rec.Body.Bytes(), &locations)
	if err != nil {
		t.Error(err.Error())
	}

	return locations
}
