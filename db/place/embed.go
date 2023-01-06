// Package place is the database of places.
package place

import (
	_ "embed"
	"encoding/json"
)

//go:embed place.json
var places string

// Places is a struct containing all the places.
type Places []string

// GetPlaces returns the Places struct.
func GetPlaces() Places {
	var p Places
	_ = json.Unmarshal([]byte(places), &p)
	return p
}
