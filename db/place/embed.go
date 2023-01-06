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

// Package place is the database of places.
package place

import (
	_ "embed"
	"encoding/json"
)

//go:embed place.json
var places string

// Places is a struct containing all the places.
type Places []string

// GetPlaces returns the Places struct.
func GetPlaces() Places {
	var p Places
	_ = json.Unmarshal([]byte(places), &p)
	return p
}
