package main

import "fmt"

func main() {
	// Deklarasi array dengan ukuran 5
	var arr [5]int

	// Operasi: Insert (inisialisasi nilai)
	arr = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array setelah insert:", arr)

	// Operasi: Akses elemen
	fmt.Println("Elemen Pada indeks ke-2 adalah", arr[2])

	// Operasi: Update elemen
	arr[0] = 24
	fmt.Println("Array setelah update:", arr)

	// Operasi: Delete elemen (dengan membuat array baru)
	indexToDelete := 1
	newArr := removeElement(arr, indexToDelete)
	fmt.Println(newArr)

	a := 2
	a = a + 5
	a = 1
	fmt.Println(a)

	//check find func
	arr2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Indeks elemen 30:", findElement(arr2, 30))   // Output: 2
	fmt.Println("Indeks elemen 100:", findElement(arr2, 100)) // Output: -1

	//check array dynamic
	var dinamis []int
	dynamicArray(&dinamis, 10)
	dynamicArray(&dinamis, 20)
	dynamicArray(&dinamis, 30)
	dynamicArray(&dinamis, 50)
	dynamicArray(&dinamis, 60)
	// fafa := dynamicArray(&dinamis, 60)
	// fmt.Println("isinya", fafa)
	fmt.Println("Array dinamis:", dinamis)

	arr3 := []int{10, 20, 30, 40, 50}
	fmt.Println("Jumlah elemen:", countElements(arr3)) // Output: 5

	arr4 := []int{10, 20, 30, 40, 50}
	fmt.Println("Elemen terbesar:", findMax(arr4)) // Output: 50
	fmt.Println("Elemen terkecil:", findMin(arr4)) // Output: 10

	arr5 := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array sebelum dibalik:", arr5)
	reversed := reverseArray1(&arr5)
	fmt.Println("Array setelah dibalik:", reversed)
	// Output:
	// Array sebelum dibalik: [10 20 30 40 50]
	// Array setelah dibalik: [50 40 30 20 10]

	arr6 := [5]int{1, 2, 3, 4, 5}
	reverseArray2(&arr6) // Mengirimkan pointer ke array
	fmt.Println(arr6)    // Output: [5 4 3 2 1]
}

// Fungsi untuk menghapus elemen dari array
func removeElement(arr [5]int, index int) []int {
	//koversi array ke slice
	slice := arr[:]
	// Hapus elemen dengan menggabungkan slice sebelum dan setelah indeks
	return append(slice[:index], slice[index+1:]...)
}

// Latihan 1: Cari Elemen dalam Array
// Buat fungsi untuk mencari elemen dalam array dan mengembalikan indeksnya. Jika tidak ditemukan, kembalikan -1.
func findElement(arr [5]int, target int) int {
	for i, value := range arr {
		if value == target {
			return i
		}
	}
	return -1
}

// Latihan 2: Implementasi Array Dinamis
// Buat program untuk menambahkan elemen ke dalam array secara dinamis menggunakan slice.
func dynamicArray(arr *[]int, element int) string {
	*arr = append(*arr, element)
	return fmt.Sprintln("Array dinamis:", *arr)
}

// Latihan 3: Menghitung Jumlah Elemen dalam Array
// Fungsi ini akan mengembalikan jumlah elemen dalam array. Karena array di Golang memiliki ukuran tetap, kita bisa menggunakan fungsi len()
func countElements(arr []int) int {
	return len(arr)
}

// Mencari Elemen Terbesar/Terkecil
// Kita akan membuat dua fungsi: satu untuk mencari elemen terbesar dan satu untuk mencari elemen terkecil.

func findMax(arr []int) int {
	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func findMin(arr []int) int {
	min := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}

// Membalik Urutan Array
// Kita akan membuat fungsi untuk membalik urutan elemen dalam array. Karena array di Golang memiliki ukuran tetap, kita perlu membuat array baru untuk menyimpan hasilnya.
func reverseArray1(arr *[5]int) [5]int {
	reverse := [5]int{} // Buat array baru untuk menyimpan hasil
	for i, value := range arr {
		reverse[len(arr)-1-i] = value // Isi array baru dari belakang
	}
	return reverse
}
func reverseArray2(arr *[5]int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
