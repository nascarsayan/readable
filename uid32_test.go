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

package readable_test

import (
	"testing"

	guid "github.com/google/uuid"
	"github.com/nascarsayan/readable"
)

func TestHashVerifier(t *testing.T) {
	for i := 0; i < 1000; i++ {
		uuid := guid.New()
		sentence, err := readable.FromUUID(uuid)
		if err != nil {
			t.Errorf("error converting uuid to sentence: %v", err)
		}
		smol, err := readable.SmolFromUUID(uuid)
		if err != nil {
			t.Errorf("error converting uuid to smol: %v", err)
		}
		if !readable.VerifySmolWithUUID(smol, uuid) {
			t.Errorf("smol is not of uuid. Smol: %v, UUID: %v, Readable: %v", smol, uuid, sentence)
		}
		if !readable.VerifySmolWithReadable(smol, sentence) {
			t.Errorf("smol is not of readable. Smol: %v, UUID: %v, Readable: %v", smol, uuid, sentence)
		}
	}
}

func TestNewSmol(t *testing.T) {
	sentence, err := readable.New()
	if err != nil {
		t.Errorf("error creating new readable: %v", err)
	}
	smol, err := readable.Smol(sentence)
	if err != nil {
		t.Errorf("error converting readable to smol: %v", err)
	}
	if len(smol) == 0 {
		t.Errorf("sentence is empty")
	}
}

func TestNewSmolFromUUID(t *testing.T) {
	smol, err := readable.SmolFromUUID(guid.New())
	if err != nil {
		t.Errorf("error converting uuid to smol: %v", err)
	}
	if len(smol) == 0 {
		t.Errorf("sentence is empty")
	}
}
