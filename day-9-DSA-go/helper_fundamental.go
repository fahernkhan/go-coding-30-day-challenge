package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	// Contoh slice
	data := []int{8, 3, 5, 1, 9, 6}
	fmt.Println("Original:", data)

	// ✅ 1. Sorting Ascending
	slices.Sort(data)
	fmt.Println("Sorted Asc:", data)

	// ✅ 2. Sorting Descending
	slices.Reverse(data)
	fmt.Println("Sorted Desc:", data)

	// ✅ 3. Reverse kembali (balik urutan)
	slices.Reverse(data)
	fmt.Println("Reversed Again:", data)

	// ✅ 4. Min dan Max dari slice
	min := slices.Min(data) // butuh Go 1.21+
	max := slices.Max(data)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	// ✅ 5. Sum / Penjumlahan total isi slice
	var sum int
	for _, v := range data {
		sum += v
	}
	fmt.Println("Sum:", sum)

	// ✅ 6. Clone (copy slice)
	cloned := slices.Clone(data)
	fmt.Println("Cloned:", cloned)

	// ✅ 7. Cek kesamaan dua slice
	isEqual := slices.Equal(data, cloned)
	fmt.Println("Equal dengan cloned:", isEqual)

	// ✅ 8. Binary Search (harus diurutkan ascending dulu)
	slices.Sort(data)
	idx, found := slices.BinarySearch(data, 5)
	if found {
		fmt.Printf("Angka 5 ditemukan di index %d\n", idx)
	} else {
		fmt.Println("Angka 5 tidak ditemukan")
	}

	// ✅ 9. Custom Sort (contoh: urutkan genap dulu, baru ganjil, lalu berdasarkan nilai)
	slices.SortFunc(data, func(a, b int) int {
		// Genap didahulukan
		if a%2 == 0 && b%2 != 0 {
			return -1
		}
		if a%2 != 0 && b%2 == 0 {
			return 1
		}
		return cmp.Compare(a, b)
	})
	fmt.Println("Custom Sort (Even First):", data)
}
