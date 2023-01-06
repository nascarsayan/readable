// Package animal is the database of animals.
package animal

import (
	_ "embed"
	"encoding/json"
)

//go:embed adjective.json
var adjectives string

//go:embed noun.json
var nouns string

//go:embed verb.json
var verbs string

// Animals is a struct containing all the animals.
type Animals struct {
	Adjective []string
	Noun      []string
	Verb      []string
}

// GetAnimals returns the Animals struct.
func GetAnimals() Animals {
	var animals Animals
	_ = json.Unmarshal([]byte(adjectives), &animals.Adjective)
	_ = json.Unmarshal([]byte(nouns), &animals.Noun)
	_ = json.Unmarshal([]byte(verbs), &animals.Verb)
	return animals
}
