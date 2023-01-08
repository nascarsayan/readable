// Licensed to Sayan Naskar under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Sayan Naskar licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Package db contains all the data for the readable package.
package db

import (
	"github.com/nascarsayan/readable/db/animal"
	"github.com/nascarsayan/readable/db/grammar"
	"github.com/nascarsayan/readable/db/name"
	"github.com/nascarsayan/readable/db/place"
)

type db struct {
	Animals animal.Animals
	Grammar grammar.Grammars
	Name    name.Names
	Place   place.Places
}

// Data is the database for the readable package.
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
