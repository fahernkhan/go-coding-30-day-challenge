package main

import "fmt"

func prefixSumSubarrayCount(arr []int, k int) int {
	prefix := 0
	count := 0
	m := make(map[int]int)
	m[0] = 1

	for _, num := range arr {
		prefix += num
		count += m[prefix-k]
		m[prefix]++
	}

	return count
}

func main() {
	arr := []int{1, 2, 3, -2, 2, 1}
	k := 3
	fmt.Printf("Jumlah subarray dengan jumlah %d: %d\n", k, prefixSumSubarrayCount(arr, k))
}
