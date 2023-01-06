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

// Package name is the database of name.
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

// Names is a struct containing all the names.
type Names struct {
	First  []string
	Last   []string
	Middle []string
}

// GetNames returns the Names struct.
func GetNames() Names {
	var names Names
	_ = json.Unmarshal([]byte(firsts), &names.First)
	_ = json.Unmarshal([]byte(middles), &names.Middle)
	_ = json.Unmarshal([]byte(lasts), &names.Last)
	return names
}
