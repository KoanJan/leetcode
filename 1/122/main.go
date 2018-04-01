package main

import (
	"fmt"
)

func maxProfit(prices []int) int {

	// no transactions
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	var (
		hold   bool = false
		i      int  = 0
		next   int  = 1
		length int  = len(prices)
		profit int  = 0
	)

	for next < length {

		// loop
		if hold {
			if prices[next] < prices[i] {
				// sell
				profit += prices[i]
				hold = false
			}
		} else {
			if prices[next] > prices[i] {
				// purchase
				profit -= prices[i]
				hold = true
			}
		}
		i += 1
		next += 1
	}

	// last day
	if hold {
		// sell always
		profit += prices[i]
		hold = false
	}

	return profit

}

func main() {

	// test
	prices := []int{1, 2}
	fmt.Printf("MaxProfit is %d while prices day by day are %v\n", maxProfit(prices), prices)
}
