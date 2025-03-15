package main

import (
	"fmt"
	"sort"
)

func main() {
	// ================================
	// 1. OPERASI DASAR PADA ARRAY & SLICE
	// ================================

	// Deklarasi array dengan ukuran tetap (fixed size)
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

	// Deklarasi slice (dinamis)
	slice := []int{1, 2, 3, 4}
	fmt.Println("Slice awal:", slice)

	// Akses elemen dalam slice
	fmt.Println("Elemen pertama:", slice[0])

	// Update elemen dalam slice
	slice[1] = 10
	fmt.Println("Slice setelah update:", slice)

	// Insert elemen ke dalam slice (menggunakan append)
	slice = append(slice, 5)
	fmt.Println("Slice setelah insert:", slice)

	// Delete elemen terakhir dari slice
	slice = slice[:len(slice)-1]
	fmt.Println("Slice setelah delete terakhir:", slice)

	// Insert elemen di posisi tertentu (misal indeks 2)
	index := 2
	slice = append(slice[:index], append([]int{99}, slice[index:]...)...)
	fmt.Println("Slice setelah insert di indeks 2:", slice)

	// Delete elemen di posisi tertentu (misal indeks 2)
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println("Slice setelah delete di indeks 2:", slice)

	// Sorting slice
	sort.Ints(slice)
	fmt.Println("Slice setelah sorting:", slice)

	// ================================
	// 2. TWO-POINTER TECHNIQUE: Two Sum (LeetCode #1)
	// ================================
	// Diberikan array nums dan target, cari dua angka yang jumlahnya sama dengan target.

	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println("Two Sum result:", result)

	// ================================
	// 3. SLIDING WINDOW: Best Time to Buy and Sell Stock (LeetCode #121)
	// ================================
	// Diberikan harga saham setiap hari, cari keuntungan maksimal yang bisa diperoleh.

	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit(prices)
	fmt.Println("Max profit:", profit)

	// ================================
	// 4. SLIDING WINDOW: Maximum Subarray Sum (Kadane's Algorithm)
	// ================================
	// Diberikan array nums, cari subarray dengan jumlah maksimum.

	nums2 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum := maxSubArray(nums2)
	fmt.Println("Maximum subarray sum:", maxSum)
}

// Fungsi Two Sum dengan pendekatan hashmap (lebih optimal daripada brute force)
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int) // Menyimpan nilai dan indeksnya
	for i, num := range nums {
		complement := target - num // Cari angka yang jika dijumlahkan dengan num menghasilkan target
		if index, found := hashMap[complement]; found {
			return []int{index, i} // Jika ditemukan, kembalikan indeks pasangan angka
		}
		hashMap[num] = i // Simpan angka dan indeksnya dalam hashmap
	}
	return []int{} // Jika tidak ditemukan, kembalikan slice kosong
}

// Fungsi maxProfit menggunakan teknik sliding window (memanfaatkan pointer minimum harga)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0] // Inisialisasi harga minimum
	maxProfit := 0        // Inisialisasi keuntungan maksimal
	for _, price := range prices {
		if price < minPrice {
			minPrice = price // Perbarui harga minimum jika ditemukan harga yang lebih rendah
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice // Hitung keuntungan maksimal jika harga saat ini lebih tinggi
		}
	}
	return maxProfit
}

// Fungsi maxSubArray menggunakan Kadane's Algorithm untuk mencari subarray dengan jumlah maksimum
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

// Fungsi helper untuk mencari nilai maksimum dari dua angka
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
