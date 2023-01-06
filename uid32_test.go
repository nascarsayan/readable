package fluuid_test

import (
	"testing"

	guid "github.com/google/uuid"
	"github.com/nascarsayan/fluuid"
)

func TestHashVerifier(t *testing.T) {
	for i := 0; i < 3; i++ {
		uuid := guid.New()
		sentence, err := fluuid.FromUUID(uuid)
		if err != nil {
			t.Errorf("error converting uuid to sentence: %v", err)
		}
		smol, err := fluuid.SmolFromUUID(uuid)
		if err != nil {
			t.Errorf("error converting uuid to smol: %v", err)
		}
		if !fluuid.VerifySmolWithUUID(smol, uuid) {
			t.Errorf("smol is not of uuid. Smol: %v, UUID: %v Fluuid: %v", smol, uuid, sentence)
		}
		if !fluuid.VerifySmolWithFluuid(smol, sentence) {
			t.Errorf("smol is not of fluuid. Smol: %v, UUID: %v Fluuid: %v", smol, uuid, sentence)
		}
	}
}

func TestNewSmol(t *testing.T) {
	sentence, err := fluuid.New()
	if err != nil {
		t.Errorf("error creating new smol: %v", err)
	}
	smol, err := fluuid.Smol(sentence)
	if err != nil {
		t.Errorf("error creating new smol: %v", err)
	}
	if len(smol) == 0 {
		t.Errorf("sentence is empty")
	}
}

func TestNewSmolFromUUID(t *testing.T) {
	smol, err := fluuid.SmolFromUUID(guid.New())
	if err != nil {
		t.Errorf("error creating new smol: %v", err)
	}
	if len(smol) == 0 {
		t.Errorf("sentence is empty")
	}
}
