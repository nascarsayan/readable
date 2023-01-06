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
