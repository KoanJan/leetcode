package main

import "fmt"

func isPalindrome(s string) bool {
	data := []byte(s)
	ln := len(data)
	i, j := 0, ln-1
	for i <= j {
		for ; i < ln; i++ {
			if 65 <= data[i] && data[i] <= 90 {
				data[i] += 32
				break
			} else if (97 <= data[i] && data[i] <= 122) || (48 <= data[i] && data[i] <= 57) {
				break
			}
		}
		for ; 0 <= j; j-- {
			if 65 <= data[j] && data[j] <= 90 {
				data[j] += 32
				break
			} else if (97 <= data[j] && data[j] <= 122) || (48 <= data[j] && data[j] <= 57) {
				break
			}
		}
		if i < ln && j >= 0 && data[i] != data[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(int('a')) // 97
	fmt.Println(int('z')) // 122
	fmt.Println(int('A')) // 65
	fmt.Println(int('Z')) // 90
	fmt.Println(int('0')) // 48
	fmt.Println(int('9')) // 57

	//s := "A man, a plan, a canal: Panama"
	s := ".,"
	fmt.Println(isPalindrome(s))
}
