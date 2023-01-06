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

package fluuid

import (
	"github.com/google/uuid"
	"github.com/nascarsayan/fluuid/db"
)

var uid128 *Converter

// InitUID128 initializes the fluuid package.
func InitUID128() {
	schema := [][]string{
		db.Data.Name.First,
		db.Data.Name.Middle,
		db.Data.Name.Last,
		{"the"},
		db.Data.Grammar.PersonalNoun,
		{"of"},
		db.Data.Place,
		db.Data.Grammar.Verb,
		db.Data.Name.First,
		db.Data.Name.Middle,
		db.Data.Name.Last,
		{"and"},
		getRange(2, 34),
		db.Data.Animals.Adjective,
		db.Data.Animals.Noun,
	}
	bitSizeList := computeBitSizeList(schema)
	// bitSizeList = []int{12, 11, 14, 0, 13, 0, 13, 10, 12, 11, 14, 0, 5, 6, 7}
	uid128 = &Converter{&schema, &bitSizeList}
}

// New returns a new fluuid.
func New() (string, error) {
	u := uuid.New()
	b, err := u.MarshalBinary()
	if err != nil {
		return "", err
	}
	return uid128.Marshal(b)
}

// ToUUID converts a fluuid to a uuid.
func ToUUID(fluuid string) (*uuid.UUID, error) {
	b := make([]byte, 16)
	err := uid128.Unmarshal(fluuid, &b)
	if err != nil {
		return nil, err
	}
	u := new(uuid.UUID)
	err = u.UnmarshalBinary(b)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FromUUID converts a uuid to a fluuid.
func FromUUID(u uuid.UUID) (string, error) {
	return uid128.Marshal(u[:])
}
