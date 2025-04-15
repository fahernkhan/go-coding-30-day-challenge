package main

import "fmt"

func twoSum(nums []int, target int) (int, int) {
	seen := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if idx, ok := seen[diff]; ok {
			return idx, i
		}
		seen[num] = i
	}
	return -1, -1
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	i, j := twoSum(nums, target)
	if i != -1 {
		fmt.Printf("Pasangan ditemukan di index %d dan %d\n", i, j)
	} else {
		fmt.Println("Tidak ada pasangan.")
	}
}
