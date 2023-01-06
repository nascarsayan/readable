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
