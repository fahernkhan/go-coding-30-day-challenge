package main

import (
	"fmt"
	"sort"
)

// printIndex prints the index of the searched number
func printIndex(numbers []int, target int) {
	index := sort.SearchInts(numbers, target)
	fmt.Printf("Angka %d ditemukan pada indeks %d\n", target, index)
}
func main() {
	// Search for the index of an element in a sorted slice of integers using binary search
	// Data slice yang sudah terurut
	numbers := []int{1, 2, 3, 4, 5, 6}

	// Cari angka tertentu dalam slice
	target := 4

	// Cetak hasil pencarian
	printIndex(numbers, target)
}
