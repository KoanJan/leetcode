package main

import "fmt"

func reverseWords(s string) string {
	data := []byte(s)
	ln := len(data)
	i, j := 0, 0
	for i < ln {
		for j < ln {
			if data[j] == ' ' {
				break
			}
			j++
		}
		for k := i; k <= (i+j-1)/2; k++ {
			data[k], data[j-1-k+i] = data[j-1-k+i], data[k]
		}
		i = j + 1
		j = i
	}
	return string(data)
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}
