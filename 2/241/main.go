package main

import (
	"fmt"
)

func diffWaysToCompute(input string) []int {
	return dw2c(formula(input))
}

func dw2c(fml []int) []int {
	printFormula(fml)
	// it's one result
	if len(fml) == 1 {
		return fml
	}
	if len(fml)%2 != 1 {
		panic("invalid formula")
	}
	var (
		newFml = []int{}
		res    = []int{}
	)
	for i := 0; i < len(fml)-2; i += 2 {
		for j := 0; j < i; j += 1 {
			newFml = append(newFml, fml[j])
		}
		newFml = append(newFml, calc(fml[i:i+3]))
		for j := i + 3; j < len(fml); j += 1 {
			newFml = append(newFml, fml[j])
		}
		tmp := dw2c(newFml)
		for _, r := range tmp {
			res = append(res, r)
		}
		newFml = []int{}
	}
	return res
}

func formula(input string) []int {
	var (
		bi      = []byte(input)
		res     = []int{}
		l       = len(bi)
		i   int = 0
		cur     = []byte{}
	)
	for i < l {
		if isSymbol(bi[i]) {
			if len(cur) == 0 {
				panic("invalid input")
			}
			res = append(res, number(cur))
			cur = []byte{}
			res = append(res, int(bi[i]))
		} else {
			cur = append(cur, bi[i])
		}
		i += 1
	}
	if len(cur) == 0 {
		panic("invalid input")
	}
	res = append(res, number(cur))
	return res
}

func isSymbol(c byte) bool {
	return c == '*' || c == '-' || c == '+'
}

func number(b []byte) int {
	var (
		length     = len(b)
		d      int = 1
		res    int = 0
	)
	for i := length - 1; i >= 0; i-- {
		res += (int(b[i]) - 48) * d
		d *= 10
	}
	return res
}

func calc(fml []int) int {
	if len(fml) == 1 {
		return fml[0]
	}
	if len(fml) != 3 {
		panic("invalid formula")
	}
	switch fml[1] {
	case int('*'):
		return fml[0] * fml[2]
	case int('+'):
		return fml[0] + fml[2]
	case int('-'):
		return fml[0] - fml[2]
	default:
		panic("invalid operator")
	}
}

func printFormula(fml []int) {
	for _, e := range fml {
		if isSymbol(byte(e)) {
			fmt.Printf("%c ", e)
		} else {
			fmt.Printf("%d ", e)
		}
	}
	fmt.Printf("\n")
}

func main() {
	// test
	input := "2*3-4*5"
	fmt.Printf("input: %s\noutput: %v\n", input, diffWaysToCompute(input))
}
