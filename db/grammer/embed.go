package grammer

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

type Grammers struct {
	Adjective    []string
	Adverb       []string
	PersonalNoun []string
	Verb         []string
}

func GetGrammers() Grammers {
	var grammars Grammers
	_ = json.Unmarshal([]byte(adjectives), &grammars.Adjective)
	_ = json.Unmarshal([]byte(adverbs), &grammars.Adverb)
	_ = json.Unmarshal([]byte(personalNouns), &grammars.PersonalNoun)
	_ = json.Unmarshal([]byte(verbs), &grammars.Verb)
	return grammars
}
