package fluuid

import (
	"fmt"
	"math"
	"strings"

	"github.com/google/uuid"
	"github.com/nascarsayan/fluuid/db"
)

var schema [][]string
var bitSizeList []int

func init() {
	db.InitDB()
	initBitSizeList()
}

func Init() {
	db.InitDB()
	initBitSizeList()
}

func getRange(start int, end int) []string {
	r := make([]string, end-start)
	for i := start; i < end; i++ {
		r[i-start] = fmt.Sprintf("%d", i)
	}
	return r
}

func initBitSizeList() {
	schema = [][]string{
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
	bitSizeList = make([]int, len(schema))
	for i, v := range schema {
		bitSizeList[i] = int(math.Log2(float64(len(v))))
	}
	bitSizeTotal := 0
	for _, v := range bitSizeList {
		bitSizeTotal += v
	}
	if bitSizeTotal < 128 {
		panic("bitSizeTotal < 128")
	}
}

type ReadableUUID struct {
	GUID uuid.UUID
}

func New() *ReadableUUID {
	return &ReadableUUID{
		GUID: uuid.New(),
	}
}

func (u *ReadableUUID) String() string {
	return FromUUID(u.GUID)
}

func FromUUID(uuid uuid.UUID) string {
	res := ""
	bitIdx, byteIdx := 0, 0

	nextChunk := func(bitSize int) int {
		r := 0
		for i := 0; i < bitSize; i++ {
			r = r << 1
			r = r | (int(uuid[byteIdx]) >> (7 - bitIdx) & 1)
			bitIdx++
			if bitIdx == 8 {
				byteIdx++
				bitIdx = 0
			}
		}
		return r
	}

	for i := 0; i < len(bitSizeList); i++ {
		res += schema[i][nextChunk(bitSizeList[i])]
		if i != len(bitSizeList)-1 {
			res += " "
		}
	}
	return res
}

func indexOf(s *[]string, v string) int {
	for i, x := range *s {
		if x == v {
			return i
		}
	}
	return -1
}

func ToUUID(readable string) (uuid.UUID, error) {
	res := make([]byte, 16)
	bitIdx, byteIdx, currentByte := 0, 0, 0
	words := strings.Split(readable, " ")
	for i := 0; i < len(bitSizeList); i++ {
		idx := indexOf(&schema[i], words[i])
		if idx == -1 {
			return uuid.Nil, fmt.Errorf("invalid word: %s", words[i])
		}
		for j := 0; j < bitSizeList[i]; j++ {
			currentByte = currentByte << 1
			currentByte = currentByte | (idx >> (bitSizeList[i] - j - 1) & 1)
			bitIdx++
			if bitIdx == 8 {
				res[byteIdx] = byte(currentByte)
				byteIdx++
				bitIdx = 0
				currentByte = 0
			}
		}
	}
	return uuid.FromBytes(res)
}

func testUuid2Readable() {
	readable := New()
	fmt.Printf("%v\n", readable.GUID)
	fmt.Printf("%v\n", readable)
	uuid, err := ToUUID(readable.String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", uuid)
	uuid2 := FromUUID(uuid)
	fmt.Printf("%v\n", uuid2)
}

func Test() {
	for i := 0; i < 100; i++ {
		testUuid2Readable()
	}
}
