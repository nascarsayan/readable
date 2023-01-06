package fluuid

import (
	"fmt"
	"math"

	"github.com/nascarsayan/fluuid/db"
)

func init() {
	db.InitDB()
	InitUID128()
	InitUID32()
}

func indexOf(s *[]string, v string) int {
	for i, x := range *s {
		if x == v {
			return i
		}
	}
	return -1
}

func getRange(start int, end int) []string {
	r := make([]string, end-start)
	for i := start; i < end; i++ {
		r[i-start] = fmt.Sprintf("%d", i)
	}
	return r
}

func computeBitSizeList(schema [][]string) []int {
	bitSizeList := make([]int, len(schema))
	for i, v := range schema {
		bitSizeList[i] = int(math.Log2(float64(len(v))))
	}
	return bitSizeList
}
