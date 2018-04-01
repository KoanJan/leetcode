package main

import "fmt"

func convert(s string, numRows int) string {
	if s == "" || numRows <= 1 {
		return s
	}
	ln := len(s)
	if ln <= numRows {
		return s
	}
	src := []byte(s)
	result := make([]byte, 0, ln)
	for i := 0; i < ln; i += 2 * (numRows - 1) {
		result = append(result, src[i])
	}
	for r := 1; r < numRows-1; r++ {
		for i := r; i < ln; {
			result = append(result, src[i])
			innerIndex := i + 2*(numRows-r-1)
			if innerIndex >= ln {
				break
			}
			result = append(result, src[innerIndex])
			i += 2 * (numRows - 1)
			if i >= ln {
				break
			}
		}
	}
	for i := numRows - 1; i < ln; i += 2 * (numRows - 1) {
		result = append(result, src[i])
	}
	return string(result)
}

func main() {
	//s := "PAYPALISHIRING"
	s := "AB"
	s2 := convert(s, 1)
	fmt.Println(s2)
}
