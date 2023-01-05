package fluuid_test

import (
	"testing"

	"github.com/nascarsayan/fluuid"
)

func TestGetReadableFromUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		f := fluuid.New()
		sentence := f.String()
		uuid := f.GUID
		uuidExpected, err := fluuid.ToUUID(sentence)
		if err != nil {
			t.Errorf("error converting sentence to uuid: %v", err)
		}
		if uuid != uuidExpected {
			t.Errorf("uuids are not equal: %v != %v. Sentence is: %v", uuid, uuidExpected, sentence)
		}
		sentenceExpected := fluuid.FromUUID(uuidExpected)
		if sentence != sentenceExpected {
			t.Errorf("sentences are not equal: %v != %v. UUID is: %v", sentence, sentenceExpected, uuid)
		}
	}
}
