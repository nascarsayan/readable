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

// Package grammar is the database of grammar.
package grammar

import (
	_ "embed"
	"encoding/json"
)

//go:embed adjective.json
var adjectives string

//go:embed adverb.json
var adverbs string

//go:embed personal-noun.json
var personalNouns string

//go:embed verb.json
var verbs string

// Grammars is a struct containing all the grammars.
type Grammars struct {
	Adjective    []string
	Adverb       []string
	PersonalNoun []string
	Verb         []string
}

// GetGrammers returns the Grammars struct.
func GetGrammers() Grammars {
	var grammars Grammars
	_ = json.Unmarshal([]byte(adjectives), &grammars.Adjective)
	_ = json.Unmarshal([]byte(adverbs), &grammars.Adverb)
	_ = json.Unmarshal([]byte(personalNouns), &grammars.PersonalNoun)
	_ = json.Unmarshal([]byte(verbs), &grammars.Verb)
	return grammars
}
