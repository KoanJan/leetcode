package main

import (
	"fmt"
)

// Symbol	I	V	X	L	C	D	M
// Value	1	5	10	50	100	500	1,000
func intToRoman(num int) string {
	var (
		res string = ""
	)

	// M
	if num >= 1000 {
		for num >= 1000 {
			num -= 1000
			res += "M"
		}
	}

	// CM
	if num >= 900 {
		num -= 900
		res += "CM"
	}

	// D
	if num >= 500 {
		for num >= 500 {
			res += "D"
			num -= 500
		}
	}

	// CD
	if num >= 400 {
		num -= 400
		res += "CD"
	}

	// C
	if num >= 100 {
		for num >= 100 {
			res += "C"
			num -= 100
		}
	}

	// XC
	if num >= 90 {
		num -= 90
		res += "XC"
	}

	// L
	if num >= 50 {
		for num >= 50 {
			res += "L"
			num -= 50
		}
	}

	// XL
	if num >= 40 {
		num -= 40
		res += "XL"
	}

	// X
	if num >= 10 {
		for num >= 10 {
			res += "X"
			num -= 10
		}
	}

	// IX
	if num >= 9 {
		num -= 9
		res += "IX"
	}

	// V
	if num >= 5 {
		for num >= 5 {
			res += "V"
			num -= 5
		}
	}

	// IV
	if num >= 4 {
		num -= 4
		res += "IV"
	}

	// I
	for num >= 1 {
		res += "I"
		num -= 1
	}

	return res
}

func main() {

	// test
	num := 3999
	fmt.Printf("intToRoman(%d) = %s", num, intToRoman(num))
}
