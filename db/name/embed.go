package name

import (
	_ "embed"
	"encoding/json"
)

//go:embed first.json
var firsts string

//go:embed last.json
var lasts string

//go:embed middle.json
var middles string

type Names struct {
	First  []string
	Last   []string
	Middle []string
}

func GetNames() Names {
	var names Names
	_ = json.Unmarshal([]byte(firsts), &names.First)
	_ = json.Unmarshal([]byte(middles), &names.Middle)
	_ = json.Unmarshal([]byte(lasts), &names.Last)
	return names
}
