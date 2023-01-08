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

package readable

import (
	"hash/fnv"

	"github.com/google/uuid"
	"github.com/nascarsayan/readable/db"
)

var uid32 *Converter

// InitUID32 initializes the uid32 converter.
func InitUID32() {
	schema := [][]string{
		getRange(2, 66),
		db.Data.Animals.Adjective,
		db.Data.Animals.Noun,
		db.Data.Grammar.Adverb,
		db.Data.Animals.Verb,
	}
	bitSizeList := computeBitSizeList(schema)
	// bitSizeList = []int{6, 6, 7, 8, 5}
	uid32 = &Converter{&schema, &bitSizeList}
}

func createHash(u uuid.UUID) [4]byte {
	h := fnv.New32a()
	h.Write(u[:])
	hash := h.Sum32()
	var bytes [4]byte
	for i := 0; i < 4; i++ {
		bytes[i] = byte(hash >> uint(8*i) & 0xff)
	}
	return bytes
}

// Smol returns a 32-bit hash of the given readable.
func Smol(readable string) (string, error) {
	b := make([]byte, 16)
	err := uid128.Unmarshal(readable, &b)
	if err != nil {
		return "", err
	}
	u := new(uuid.UUID)
	err = u.UnmarshalBinary(b)
	if err != nil {
		return "", err
	}
	return SmolFromUUID(*u)
}

// SmolFromUUID returns a 32-bit hash of the given uuid.
func SmolFromUUID(u uuid.UUID) (string, error) {
	hash := createHash(u)
	return uid32.Marshal(hash[:])
}

// VerifySmolWithReadable verifies that the given smol is the 32-bit hash of the given readable.
func VerifySmolWithReadable(smol string, readable string) bool {
	smol2, err := Smol(readable)
	if err != nil {
		return false
	}
	return smol == smol2
}

// VerifySmolWithUUID verifies that the given smol is the 32-bit hash of the given uuid.
func VerifySmolWithUUID(smol string, u uuid.UUID) bool {
	smol2, err := SmolFromUUID(u)
	if err != nil {
		return false
	}
	return smol == smol2
}
