package main

import (
	"fmt"
)

func lastRemaining(n int) int {
	return remaining(1, n, 2, true)
}

func remaining(start, end, pace int, left bool) int {
	// fmt.Printf("start: %d end: %d pace: %d left: %t\n", start, end, pace, left)
	if end-start >= pace {
		if left {
			if (end-start+1)%pace != 1 {
				return remaining(start+pace/2, end, pace*2, false)
			} else {
				return remaining(start+pace/2, end-pace/2, pace*2, false)
			}
		} else {
			if (end-start+1)%pace != 1 {
				return remaining(start, end-pace/2, pace*2, true)
			} else {
				return remaining(start+pace/2, end-pace/2, pace*2, true)
			}
		}
	} else {
		if left {
			return end
		} else {
			return start
		}
	}
}

func main() {
	// test
	n := 2
	fmt.Printf("lastRemaining(%d) = %d\n", n, lastRemaining(n))
}
