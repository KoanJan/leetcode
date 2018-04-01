package main

import "fmt"

func originalDigits(s string) string {
	data := []byte(s)
	count := make([]int, 10)
	// collect
	for _, b := range data {
		switch b {
		case 'z': // 0
			count[0]++
		case 'x': // 6
			count[6]++
		case 's': // 7(6)
			count[7]++
		case 'u': // 4
			count[4]++
		case 'w': // 2
			count[2]++
		case 'f': // 5(4)
			count[5]++
		case 'g': // 8
			count[8]++
		case 'o': // 1(0,2,4)
			count[1]++
		case 'r': // 3(4,0)
			count[3]++
		case 'i': // 9(8,6,5)
			count[9]++
		}
	}
	// fix
	count[7] -= count[6]
	count[5] -= count[4]
	count[1] = count[1] - count[0] - count[2] - count[4]
	count[3] = count[3] - count[4] - count[0]
	count[9] = count[9] - count[8] - count[6] - count[5]

	// return
	res := []byte{}
	for i := range count {
		c := count[i]
		for j := 0; j < c; j++ {
			res = append(res, byte(i+48))
		}
	}
	return string(res)
}

func main() {
	s := "owoztneoer"
	fmt.Println(originalDigits(s))
}
