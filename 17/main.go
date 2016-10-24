package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {

	data := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
		'*': []byte{'g', 'h'},
		'#': []byte{'`'},
	}

	var (
		res = []string{}
		bs  = []byte(digits)
	)

	if len(bs) == 0 {
		return res
	}
	for i := 0; i < len(bs); i++ {
		if _, existed := data[bs[i]]; !existed {
			return res
		}
	}

	subCombinations := []string{""}
	if len(bs) > 1 {
		subCombinations = letterCombinations(string(bs[1:]))
	}

	for i := 0; i < len(data[bs[0]]); i++ {
		for j := 0; j < len(subCombinations); j++ {
			res = append(res, string([]byte{data[bs[0]][i]})+subCombinations[j])
		}
	}

	return res
}

func main() {
	// test
	digits := ""
	fmt.Printf("Letter combinations of %s is %v\n", digits, letterCombinations(digits))
}
