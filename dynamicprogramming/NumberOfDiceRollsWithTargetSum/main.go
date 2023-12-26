package main

import (
	"fmt"
)

func numRollsToTarget(d int, f int, target int) int {
	const mod = 1e9 + 7
	dp := make([][]int, d+1)

	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	dp[0][0] = 1

	for i := 1; i <= d; i++ {
		for j := 1; j <= target; j++ {
			for k := 1; k <= f && k <= j; k++ {
				dp[i][j] = (dp[i][j] + dp[i-1][j-k]) % mod
			}
		}
	}

	return dp[d][target]
}

func main() {
	d := 2
	f := 6
	target := 7

	result := numRollsToTarget(d, f, target)
	fmt.Println(result)
}
