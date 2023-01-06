package fluuid

import (
	"hash/fnv"

	"github.com/google/uuid"
	"github.com/nascarsayan/fluuid/db"
)

var uid32 *Converter

func InitUID32() {
	schema := [][]string{
		getRange(2, 66),
		db.Data.Animals.Adjective,
		db.Data.Animals.Noun,
		db.Data.Grammer.Adverb,
		db.Data.Animals.Verb,
	}
	bitSizeList := computeBitSizeList(schema)
	// bitSizeList = []int{12, 11, 14, 0, 13, 0, 13, 10, 12, 11, 14, 0, 5, 6, 7}
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

func Smol(fluuid string) (string, error) {
	b := make([]byte, 16)
	err := uid128.Unmarshal(fluuid, &b)
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

func SmolFromUUID(u uuid.UUID) (string, error) {
	hash := createHash(u)
	return uid32.Marshal(hash[:])
}

func VerifySmolWithFluuid(smol string, fluuid string) bool {
	smol2, err := Smol(fluuid)
	if err != nil {
		return false
	}
	return smol == smol2
}

func VerifySmolWithUUID(smol string, u uuid.UUID) bool {
	smol2, err := SmolFromUUID(u)
	if err != nil {
		return false
	}
	return smol == smol2
}
