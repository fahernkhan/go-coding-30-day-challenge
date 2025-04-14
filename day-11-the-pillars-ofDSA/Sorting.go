// 1. Sorting – Sort Array of Structs
// Soal: Diberikan n pelajar dengan nama dan nilai. Urutkan berdasarkan nilai menurun,
// kalau nilai sama, urutkan nama naik (alphabetical).
package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Score int
}

func main() {
	students := []Student{
		{"Andi", 90},
		{"Budi", 95},
		{"Caca", 90},
		{"Deni", 100},
	}

	// Sorting logic: Score descending, Name ascending
	sort.Slice(students, func(i, j int) bool {
		if students[i].Score == students[j].Score {
			return students[i].Name < students[j].Name
		}
		return students[i].Score > students[j].Score
	})

	fmt.Println("Hasil sorting:")
	for _, s := range students {
		fmt.Printf("%s: %d\n", s.Name, s.Score)
	}
}
