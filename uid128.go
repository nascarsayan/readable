package fluuid

import (
	"github.com/google/uuid"
	"github.com/nascarsayan/fluuid/db"
)

var uid128 *Converter

func InitUID128() {
	schema := [][]string{
		db.Data.Name.First,
		db.Data.Name.Middle,
		db.Data.Name.Last,
		{"the"},
		db.Data.Grammer.PersonalNoun,
		{"of"},
		db.Data.Place,
		db.Data.Grammer.Verb,
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

func New() (string, error) {
	u := uuid.New()
	b, err := u.MarshalBinary()
	if err != nil {
		return "", err
	}
	return uid128.Marshal(b)
}

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

func FromUUID(u *uuid.UUID) (string, error) {
	return uid128.Marshal(u[:])
}
