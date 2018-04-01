package main

import (
	"fmt"
)

func fractionToDecimal(numerator int, denominator int) string {
	if denominator == 0 {
		return ""
	}
	if numerator == 0 {
		return "0"
	}
	var (
		res            string
		resIsMinus     bool = false
		next           int
		tail           []byte = []byte{}
		tailNumerators []int  = []int{}
	)
	if numerator < 0 {
		numerator = -numerator
		resIsMinus = !resIsMinus
	}
	if denominator < 0 {
		denominator = -denominator
		resIsMinus = !resIsMinus
	}
	res += int2String(numerator / denominator)
	next = numerator % denominator
	if next == 0 {
		if resIsMinus {
			res = "-" + res
		}
		return res
	}
	res += "."
	for next != 0 {
		tailNumerators = append(tailNumerators, next)
		next *= 10
		tail = append(tail, byte(next/denominator+48))
		if yes, a := containsDup(tailNumerators); yes {
			res += string(tail[0:a]) + "(" + string(tail[a:len(tail)-1]) + ")"
			if resIsMinus {
				res = "-" + res
			}
			return res
		}
		next %= denominator
	}
	res += string(tail)
	if resIsMinus {
		res = "-" + res
	}
	return res
}

func containsDup(a []int) (bool, int) {
	var l int = len(a)
	if l < 2 {
		return false, -1
	}
	for i := l - 2; i >= 0; i-- {
		if a[i] == a[l-1] {
			return true, i
		}
	}
	return false, -1
}

func int2String(n int) string {
	if n == 0 {
		return "0"
	}
	var (
		d       []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		res     string
		isMinus bool = false
	)
	if n < 0 {
		isMinus = true
		n = -n
	}
	for n > 0 {
		res = d[n%10] + res
		n /= 10
	}
	if isMinus {
		res = "-" + res
	}
	return res
}

func main() {
	// test
	var numerator, denominator int = -2147483648, 1
	fmt.Printf("%d / %d = %s\n", numerator, denominator, fractionToDecimal(numerator, denominator))
}
