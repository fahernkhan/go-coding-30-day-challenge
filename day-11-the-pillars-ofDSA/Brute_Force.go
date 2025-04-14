// 3. Brute Force – Subset Sum
// Soal: Diberikan array bilangan dan target,
// apakah ada subset yang jumlahnya sama dengan target?

package main

import (
	"fmt"
)

func subsetSum(nums []int, target int) bool {
	n := len(nums)
	// Coba semua kombinasi (2^n kemungkinan)
	for mask := 0; mask < (1 << n); mask++ {
		sum := 0
		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 {
				sum += nums[i]
			}
		}
		if sum == target {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{3, 1, 5, 9, 12}
	target := 15

	if subsetSum(arr, target) {
		fmt.Println("Ada subset yang jumlahnya tepat.")
	} else {
		fmt.Println("Tidak ada subset yang cocok.")
	}
}
