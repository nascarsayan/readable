package place

import (
	_ "embed"
	"encoding/json"
)

//go:embed place.json
var places string

type Places []string

func GetPlaces() Places {
	var p Places
	_ = json.Unmarshal([]byte(places), &p)
	return p
}
