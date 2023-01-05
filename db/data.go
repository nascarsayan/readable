package db

import (
	"github.com/nascarsayan/fluuid/db/animal"
	"github.com/nascarsayan/fluuid/db/grammer"
	"github.com/nascarsayan/fluuid/db/name"
	"github.com/nascarsayan/fluuid/db/place"
)

type db struct {
	Animals animal.Animals
	Grammer grammer.Grammers
	Name    name.Names
	Place   place.Places
}

var Data db

func InitDB() {
	Data = db{
		Animals: animal.GetAnimals(),
		Grammer: grammer.GetGrammers(),
		Name:    name.GetNames(),
		Place:   place.GetPlaces(),
	}
}
