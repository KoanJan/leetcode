package main

import (
	"fmt"
)

func isNumber(s string) bool {
	data := trimSpace([]byte(s))
	eIndex := -1
	for i := range data {
		if data[i] == 'e' {
			if eIndex != -1 {
				return false
			}
			eIndex = i
		}
	}
	if eIndex != -1 {
		return isNativeNumber(trimSignal(data[0:eIndex])) && isNativeInteger(trimSignal(data[eIndex+1:]))
	} else {
		return isNativeNumber(trimSignal(data))
	}
}

func trimSpace(data []byte) []byte {
	//fmt.Println("trimSpace")
	//fmt.Println(string(data))
	h, t := 0, len(data)-1
	for ; h <= t; h++ {
		if data[h] != ' ' {
			break
		}
	}
	for ; t >= 0; t-- {
		if data[t] != ' ' {
			break
		}
	}
	if h > t {
		return []byte{}
	}
	return data[h : t+1]
}

func trimSignal(data []byte) []byte {
	//fmt.Println("trimSignal")
	//fmt.Println(string(data))
	if len(data) > 0 && (data[0] == '-' || data[0] == '+') {
		return data[1:]
	}
	return data
}

func isNativeNumber(data []byte) bool {
	//fmt.Println("isNativeNumber")
	//fmt.Println(string(data))
	ln := len(data)
	if ln == 0 {
		return false
	}
	hasNumber := false
	hasDot := false
	for _, b := range data {
		if b == '.' {
			if hasDot {
				return false
			}
			hasDot = true
		} else if b <= '9' && b >= '0' {
			hasNumber = true
		} else {
			return false
		}
	}
	return hasNumber
}

func isNativeInteger(data []byte) bool {
	//fmt.Println("isNativeInteger")
	//fmt.Println(string(data))
	ln := len(data)
	if ln == 0 {
		return false
	}
	for _, b := range data {
		if b > '9' || b < '0' {
			return false
		}
	}
	return true
}

func main() {
	s := " 46.e3"
	fmt.Println(isNumber(s))
}
