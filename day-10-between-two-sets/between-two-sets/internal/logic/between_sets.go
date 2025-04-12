package logic

// GetTotalX menghitung berapa banyak angka yang:
// 1. Merupakan kelipatan dari semua elemen di array a
// 2. Merupakan faktor dari semua elemen di array b
//
// Contoh:
// a = [2, 4]
// b = [16, 32, 96]
// Maka hasilnya = 3 (angka yang memenuhi: 4, 8, 16)
//
// Pendekatan: Brute-force/tabulasi dari angka 1 hingga maksimum nilai di b
func GetTotalX(a []int32, b []int32) int32 {
	// Variabel untuk menyimpan jumlah angka yang valid
	var count int32 = 0

	// Loop dari 1 sampai nilai maksimum dari array b
	// Karena angka yang bisa jadi faktor b pasti tidak lebih besar dari elemen terbesar di b
	for x := int32(1); x <= max(b); x++ {
		valid := true

		// Cek apakah x adalah kelipatan dari semua elemen di a
		// Artinya: x harus bisa dibagi habis oleh semua elemen a
		for _, val := range a {
			if x%val != 0 {
				valid = false // Jika tidak, langsung tandai tidak valid
				break
			}
		}

		// Jika x lulus syarat pertama, lanjut cek syarat kedua
		if valid {
			// Cek apakah x adalah faktor dari semua elemen di b
			// Artinya: semua elemen b harus bisa dibagi habis oleh x
			for _, val := range b {
				if val%x != 0 {
					valid = false // Jika tidak, berarti tidak valid
					break
				}
			}
		}

		// Jika x lolos kedua syarat, tambahkan ke hitungan
		if valid {
			count++
		}
	}

	// Kembalikan jumlah angka yang valid
	return count
}

// Fungsi bantu untuk mencari nilai maksimum dari array
func max(arr []int32) int32 {
	maxVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
