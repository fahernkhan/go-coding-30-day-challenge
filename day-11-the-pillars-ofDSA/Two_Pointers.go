package main

import (
	"fmt"
	"sort"
)

func countPairs(arr []int, target int) int {
	sort.Ints(arr)
	left, right := 0, len(arr)-1
	count := 0

	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			count++
			left++
			right--
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return count
}

func main() {
	arr := []int{1, 3, 2, 2, 4, 5}
	target := 6
	fmt.Printf("Jumlah pasangan dengan jumlah %d: %d\n", target, countPairs(arr, target))
}
