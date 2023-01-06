// Package db contains all the data for the fluuid package.
package db

import (
	"github.com/nascarsayan/fluuid/db/animal"
	"github.com/nascarsayan/fluuid/db/grammar"
	"github.com/nascarsayan/fluuid/db/name"
	"github.com/nascarsayan/fluuid/db/place"
)

type db struct {
	Animals animal.Animals
	Grammar grammar.Grammars
	Name    name.Names
	Place   place.Places
}

// Data is the database for the fluuid package.
var Data *db

// InitDB initializes the database.
func InitDB() {
	if Data != nil {
		return
	}
	Data = &db{
		Animals: animal.GetAnimals(),
		Grammar: grammar.GetGrammers(),
		Name:    name.GetNames(),
		Place:   place.GetPlaces(),
	}
}
