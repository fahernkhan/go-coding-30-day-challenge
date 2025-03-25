package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	low := 0             // Indeks awal
	high := len(arr) - 1 // Indeks akhir

	// Loop selama batas pencarian valid
	for low <= high {
		mid := (low + high) / 2 // Tentukan indeks tengah

		if arr[mid] == target {
			return mid // Jika ditemukan, kembalikan indeks
		} else if arr[mid] < target {
			low = mid + 1 // Jika target lebih besar, cari di kanan
		} else {
			high = mid - 1
		}
	}
	return -1 // Jika tidak ditemukan, kembalikan -1
}

func main() {
	// Data HARUS dalam kondisi terurut
	arr := []int{1, 3, 10, 15, 16, 18, 20, 24, 26, 29, 31, 41, 45, 48, 51, 86}
	target := 24

	// Panggil fungsi BinarySearch
	index := BinarySearch(arr, target)

	// Cetak hasil pencarian
	if index != -1 {
		fmt.Printf("Elemen %d ditemukan di indeks %d\n", target, index)
	} else {
		fmt.Printf("Elemen %d tidak ditemukan\n", target)
	}
}
