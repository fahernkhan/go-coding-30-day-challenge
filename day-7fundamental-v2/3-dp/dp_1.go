package main

import (
	"fmt"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	// Fibonacci
	fmt.Println("Fibonacci(10):", fib(10))

	// Coin Change
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println("Coin Change:", coinChange(coins, amount))

	// LCS
	text1 := "abcde"
	text2 := "ace"
	fmt.Println("Longest Common Subsequence:", longestCommonSubsequence(text1, text2))

	// Edit Distance
	word1 := "horse"
	word2 := "ros"
	fmt.Println("Edit Distance:", minDistance(word1, word2))
}

/* OUTPUT:
Fibonacci(10): 55
Coin Change: 3
Longest Common Subsequence: 3
Edit Distance: 3
*/
