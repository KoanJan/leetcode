package main

import (
	"fmt"
)

func restoreIpAddresses(s string) []string {
	return ria([]string{}, s)
}

func ria(pre []string, s string) []string {
	if (len(pre) == 4 && s != "") || (len(pre) < 4 && s == "") {
		return []string{}
	}
	if len(pre) == 4 && s == "" {
		return []string{pre[0] + "." + pre[1] + "." + pre[2] + "." + pre[3]}
	}
	var (
		data    []byte   = []byte(s)
		a, b, c []string = []string{}, []string{}, []string{}
	)
	a = ria(append(pre, string(data[0:1])), string(data[1:]))
	if len(data) >= 2 && data[0] != '0' {
		b = ria(append(pre, string(data[0:2])), string(data[2:]))
		for i := 0; i < len(b); i++ {
			a = append(a, b[i])
		}
		if len(data) >= 3 && data[0] != '0' && parseInteger(string(data[0:3])) <= 255 {
			c = ria(append(pre, string(data[0:3])), string(data[3:]))
			for i := 0; i < len(c); i++ {
				a = append(a, c[i])
			}
		}
	}
	return a
}

func parseInteger(val string) int {
	var (
		data    []byte = []byte(val)
		res     int    = 0
		isMinus bool   = false
		d       int    = 1
	)

	for i := len(data) - 1; i >= 0; i-- {
		if data[i] == '-' {
			isMinus = true
			break
		}
		res += (int(data[i]) - 48) * d
		d *= 10
	}
	if isMinus {
		res = -res
	}
	return res
}

func main() {
	// test
	s := "12140167147"
	fmt.Printf("s: %s\nresult: %v\n", s, restoreIpAddresses(s))
}
