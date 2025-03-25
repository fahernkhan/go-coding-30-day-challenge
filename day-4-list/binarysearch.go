package main

import (
	"fmt"
	"sort"
)

func main() {
	var elements []int
	elements = []int{1, 3, 16, 10, 45, 31, 15, 48, 26, 18, 41, 86, 29, 51, 20}
	var element int
	element = 3

	var i int

	// **Urutkan slice terlebih dahulu** menjadi := [1 3 10 15 16 18 20 26 29 31 41 45 48 51 86]
	sort.Ints(elements)

	// Binary search dengan sort.Search()
	i = sort.Search(len(elements), func(i int) bool {
		return elements[i] >= element
	})

	// Cek apakah elemen ditemukan
	if i < len(elements) && elements[i] == element {
		fmt.Printf("found element %d at index %d in %v\n", element, i, elements)
	} else {
		fmt.Printf("element %d not found in %v\n", element, elements)
	}
}
