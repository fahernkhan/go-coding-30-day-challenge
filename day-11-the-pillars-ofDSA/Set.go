package main

import "fmt"

func hasDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 3}
	if hasDuplicate(arr) {
		fmt.Println("Ada duplikat!")
	} else {
		fmt.Println("Semua elemen unik.")
	}
}
