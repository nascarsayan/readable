package fluuid_test

import (
	"testing"

	guid "github.com/google/uuid"
	"github.com/nascarsayan/fluuid"
)

func TestInterconversion(t *testing.T) {
	for i := 0; i < 10000; i++ {
		uuid := guid.New()
		sentence, err := fluuid.FromUUID(uuid)
		if err != nil {
			t.Errorf("error converting uuid to sentence: %v", err)
		}
		uuidExpected, err := fluuid.ToUUID(sentence)
		if err != nil {
			t.Errorf("error converting sentence to uuid: %v", err)
		}
		uuidStr := uuid.String()
		_ = uuidStr
		if uuid != *uuidExpected {
			t.Errorf("uuids are not equal: %v != %v. Sentence is: %v", uuid, uuidExpected, sentence)
		}
	}
}

func TestNew(t *testing.T) {
	sentence, err := fluuid.New()
	if err != nil {
		t.Errorf("error creating new fluuid: %v", err)
	}
	if len(sentence) == 0 {
		t.Errorf("sentence is empty")
	}
}
