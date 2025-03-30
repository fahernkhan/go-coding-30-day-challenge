package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sort a slice of integers using quicksort
	numbers := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(numbers)
	// numbers is now [1, 2, 3, 4, 5, 6]
	fmt.Println(numbers)

	// Sort a slice of strings using mergesort
	words := []string{"cherry", "apple", "blueberry"}
	sort.Strings(words)
	// words is now ["apple", "blueberry", "cherry"]
	fmt.Println(words)
}
