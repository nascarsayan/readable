// Package fluuid provides a way to convert a byte array to a fluuid string and vice versa.
package fluuid

import (
	"fmt"
	"strings"
)

// Converter is a struct that converts a byte array to a fluuid string and vice versa.
type Converter struct {
	schema      *[][]string
	bitSizeList *[]int
}

// Marshal converts a byte array to a fluuid string.
func (s *Converter) Marshal(data []byte) (string, error) {
	res := ""
	bitIdx, byteIdx := 0, 0
	nextChunk := func(bitSize int) int {
		r := 0
		for i := 0; i < bitSize; i++ {
			r <<= 1
			r |= (int(data[byteIdx]) >> (7 - bitIdx) & 1)
			bitIdx++
			if bitIdx == 8 {
				byteIdx++
				bitIdx = 0
			}
		}
		return r
	}

	for i := 0; i < len(*s.schema); i++ {
		res += (*s.schema)[i][nextChunk((*s.bitSizeList)[i])]
		if i != len(*s.bitSizeList)-1 {
			res += " "
		}
	}
	if byteIdx != len(data) {
		return "", fmt.Errorf("data too large to convert to fluuid")
	}
	return res, nil
}

// Unmarshal converts a fluuid string to a byte array.
func (s *Converter) Unmarshal(sent string, bytes *[]byte) error {
	if s == nil || s.schema == nil || s.bitSizeList == nil {
		return fmt.Errorf("Converter not initialized properly")
	}
	if len(*s.schema) != len(*s.bitSizeList) {
		return fmt.Errorf("Converter not initialized properly")
	}
	bitIdx, byteIdx, currentByte := 0, 0, 0
	words := strings.Split(sent, " ")
	for i := 0; i < len((*s.bitSizeList)); i++ {
		idx := indexOf(&(*s.schema)[i], words[i])
		if idx == -1 {
			return fmt.Errorf("word not found in database: %s", words[i])
		}
		for j := 0; j < (*s.bitSizeList)[i]; j++ {
			currentByte <<= currentByte
			currentByte |= idx >> ((*s.bitSizeList)[i] - j - 1) & 1
			bitIdx++
			if bitIdx == 8 {
				(*bytes)[byteIdx] = byte(currentByte)
				byteIdx++
				bitIdx = 0
				currentByte = 0
			}
		}
	}
	return nil
}
