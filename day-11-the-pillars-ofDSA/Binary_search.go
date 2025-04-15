package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) bool {
	index := sort.SearchInts(arr, target)
	return index < len(arr) && arr[index] == target
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 12}
	target := 7

	if binarySearch(arr, target) {
		fmt.Printf("Angka %d ditemukan!\n", target)
	} else {
		fmt.Printf("Angka %d tidak ditemukan.\n", target)
	}
}
