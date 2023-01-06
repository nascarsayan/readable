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
