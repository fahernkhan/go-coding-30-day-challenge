// data_engineer_fundamental.go

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 1. Membaca CSV
func loadCSV(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

// 2. Membersihkan data (hapus baris yang ada field kosong)
func cleanData(data [][]string) [][]string {
	var cleaned [][]string
	for i, row := range data {
		empty := false
		for _, field := range row {
			if field == "" {
				empty = true
				break
			}
		}
		if !empty || i == 0 { // selalu simpan header
			cleaned = append(cleaned, row)
		}
	}
	return cleaned
}

// 3. Transformasi (contoh: kolom "amount" dikali 1.1, hasil ke kolom "amount_usd")
func transformData(data [][]string) [][]string {
	header := append(data[0], "amount_usd")
	result := [][]string{header}

	amountIdx := -1
	for i, col := range data[0] {
		if col == "amount" {
			amountIdx = i
			break
		}
	}

	if amountIdx == -1 {
		log.Println("Kolom 'amount' tidak ditemukan, skip transformasi")
		return data
	}

	for i, row := range data {
		if i == 0 {
			continue // header sudah ditambahkan
		}
		amount, err := strconv.ParseFloat(row[amountIdx], 64)
		if err != nil {
			row = append(row, "0")
		} else {
			usd := amount * 1.1
			row = append(row, fmt.Sprintf("%.2f", usd))
		}
		result = append(result, row)
	}

	return result
}

// 4. Simpan ke CSV
func saveCSV(path string, data [][]string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.WriteAll(data)
}

// 5. Pipeline utama
func main() {
	inputPath := "data/raw_data.csv"
	outputPath := "data/cleaned_data.csv"

	data, err := loadCSV(inputPath)
	if err != nil {
		log.Fatalf("Gagal membaca file CSV: %v", err)
	}

	data = cleanData(data)
	data = transformData(data)

	err = saveCSV(outputPath, data)
	if err != nil {
		log.Fatalf("Gagal menyimpan file: %v", err)
	}

	fmt.Println("Pipeline selesai. Data disimpan di", outputPath)
}
