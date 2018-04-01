package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	kinds := len(coins)
	if kinds == 0 {
		return -1
	}
	simpleSort(coins)
	if amount < coins[0] {
		return -1
	}
	// init
	var dp []int = make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}
	dp[0] = 0
	for i := kinds - 1; i >= 0; i-- {
		if coins[i] <= amount {
			dp[coins[i]] = 1
		}
	}
	// loop
	for i := 1; i <= amount; i++ {
		min := -1
		for j := 0; j < kinds && coins[j] <= i; j++ {
			if dp[i-coins[j]] != -1 {
				if min == -1 || dp[i-coins[j]] < min {
					min = dp[i-coins[j]] + 1
				}
			}
		}
		if min != -1 {
			dp[i] = min
		}
	}
	// result
	return dp[amount]
}

func simpleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[i] > a[j] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp
			}
		}
	}
}

func main() {
	// test
	var (
		coins  []int = []int{474, 83, 404, 3}
		amount int   = 264
	)
	fmt.Printf("coins: %v\namount: %d\nchanges: %d\n", coins, amount, coinChange(coins, amount))
}

////////////////////////
// Accepted JAVA code //
////////////////////////

/*
public class Solution {
    public int coinChange(int[] coins, int amount) {
        if (amount == 0){
            return 0;
        }
        int kinds = coins.length;
        if (kinds == 0){
            return -1;
        }
        Arrays.sort(coins);
        if (amount < coins[0]){
            return -1;
        }

        // init
        int[] dp = new int[amount+1];
        Arrays.fill(dp, Integer.MAX_VALUE);
        dp[0] = 0;
        for (int i = kinds - 1; i >= 0; i--) {
            if (coins[i] <= amount) {
                dp[coins[i]] = 1;
            }
        }

        // loop
        for (int i = 1; i <= amount; i++) {
            int min = dp[i];
            for (int j = 0; j < kinds && coins[j] <= i; j++) {
                int d = dp[i-coins[j]];
                if (d != Integer.MAX_VALUE && d < min) {
                    min = d+1;
                }
            }
            dp[i] = min;
        }

        // result
        return dp[amount] == Integer.MAX_VALUE ? -1 : dp[amount];
    }
}
*/
