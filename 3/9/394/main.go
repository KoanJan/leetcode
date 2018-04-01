package main

import (
	"fmt"
)

func decodeString(s string) string {
	var (
		input    []byte = []byte(s)
		output   string = ""
		current  []byte = []byte{}
		times    []byte = []byte{}
		lenInput int    = len(input)
	)
	for i := 0; i < lenInput; {

		// times
		if isNumber(input[i]) {
			for i < lenInput {
				if isNumber(input[i]) {
					times = append(times, input[i])
					i++
				} else {
					break
				}
			}
			if input[i] != '[' {
				// normal string
				output += string(times)
				times = []byte{}
				continue
			}
		} else {
			for i < lenInput {
				if !isNumber(input[i]) {
					current = append(current, input[i])
					i++
				} else {
					break
				}
			}
			output += string(current)
			// reset current
			current = []byte{}
			continue
		}

		// current
		if input[i] == '[' {
			i++
			var (
				end int = 1
			)
			for i < lenInput {
				if input[i] == '[' {
					end++
				} else if input[i] == ']' {
					end--
				}
				if end == 0 {
					break
				}
				current = append(current, input[i])
				i++
			}
			// repeat
			c := decodeString(string(current))
			t := calcTimes(times)
			for k := 0; k < t; k++ {
				output += c
			}
			// reset current and times
			current = []byte{}
			times = []byte{}
			i++
		}
	}
	return output
}

func isNumber(b byte) bool {
	return (48 <= int(b)) && (int(b) <= 57)
}

func calcTimes(times []byte) int {
	num := 0
	for i := 0; i < len(times); i++ {
		num += (int(times[i]) - 48) * pow(10, len(times)-1-i)
	}
	return num
}

func pow(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * pow(x, y-1)
}

func main() {
	// test
	s := "2[ad]3[b]eer"
	fmt.Printf("decode info:\nsrc: %s\noutput: %s\n", s, decodeString(s))
	s = "2ad]3[b]eer"
	fmt.Printf("decode info:\nsrc: %s\noutput: %s\n", s, decodeString(s))
	s = "2[a2[a]]"
	fmt.Printf("decode info:\nsrc: %s\noutput: %s\n", s, decodeString(s))
	s = "2[20[bc]31[xy]]xd4[rt]"
	fmt.Printf("decode info:\nsrc: %s\noutput: %s\n", s, decodeString(s))
}
