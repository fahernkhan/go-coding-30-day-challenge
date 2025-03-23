package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// climbStairs menghitung jumlah cara naik ke n anak tangga menggunakan DP Bottom-Up.
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	// Inisialisasi array dp untuk menyimpan jumlah cara mencapai setiap anak tangga
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2 // Base case: 1 cara ke tangga 1, 2 cara ke tangga 2

	// Mengisi dp[i] berdasarkan jumlah cara dari dp[i-1] dan dp[i-2]
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n] // Mengembalikan jumlah cara naik ke anak tangga ke-n
}

// Fungsi utama untuk membaca input dan menampilkan output ala LeetCode
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	result := climbStairs(n)
	fmt.Println("Output:", result)
}
