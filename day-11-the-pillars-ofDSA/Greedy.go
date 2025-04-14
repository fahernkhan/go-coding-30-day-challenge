// 2. Greedy – Minimum Number of Coins
// Soal: Diberikan jumlah uang dan list pecahan koin.
// Hitung jumlah minimum koin yang dibutuhkan (pakai greedy).
package main

import (
	"fmt"
	"sort"
)

func minCoins(coins []int, amount int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(coins))) // Sort descending
	count := 0
	for _, coin := range coins {
		if amount >= coin {
			take := amount / coin
			count += take
			amount -= take * coin
		}
	}
	if amount > 0 {
		return -1 // Tidak bisa pas
	}
	return count
}

func main() {
	coins := []int{1, 5, 10, 25}
	amount := 63

	result := minCoins(coins, amount)
	if result != -1 {
		fmt.Printf("Minimum coins needed: %d\n", result)
	} else {
		fmt.Println("Tidak bisa mencapai jumlah uang dengan pecahan yang ada.")
	}
}
