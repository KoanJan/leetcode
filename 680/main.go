package main

import "fmt"

func validPalindrome(s string) bool {
	return validPalindromeBytes([]byte(s))
}

func validPalindromeBytes(data []byte) bool {
	ln := len(data)
	if ln <= 2 {
		return true
	}
	needDeletion := false
	i, j := 0, ln-1
	for i < j {
		if data[i] != data[j] {
			if needDeletion {
				return false
			}
			needDeletion = true
			return nativeValidPalindromeBytes(data[i:j]) || nativeValidPalindromeBytes(data[i+1:j+1])
		}
		i++
		j--
	}
	return true
}

func nativeValidPalindromeBytes(data []byte) bool {
	ln := len(data)
	if ln < 2 {
		return true
	}
	i, j := 0, ln-1
	for i < j {
		if data[i] != data[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	s := "abababd"
	fmt.Println(validPalindrome(s))
}
