package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// =======================
	// SLICING STRING
	// =======================
	s := "07:05:45PM"
	fmt.Println("Original string:", s)

	// left to right bound like read text
	// Ambil 2 karakter terakhir (AM atau PM)
	period := s[len(s)-2:]
	fmt.Println("Sliced period (len(s)-2:):", period)

	// Ambil bagian waktu tanpa AM/PM
	timeOnly := s[:len(s)-2]
	fmt.Println("Sliced time (s[:len(s)-2]):", timeOnly)

	// =======================
	// SPLIT STRING
	// =======================
	timeParts := strings.Split(timeOnly, ":") // hasilnya slice string
	fmt.Println("Split time into parts:", timeParts)

	hourStr := timeParts[0]
	minuteStr := timeParts[1]
	secondStr := timeParts[2]

	// =======================
	// INDEXING
	// =======================
	fmt.Println("Indexing timeParts[0] (hour):", hourStr)
	fmt.Println("Indexing timeParts[1] (minute):", minuteStr)
	fmt.Println("Indexing timeParts[2] (second):", secondStr)

	// =======================
	// CONVERT TO INT & LOGIC
	// =======================
	hour, _ := strconv.Atoi(hourStr)

	if period == "PM" && hour != 12 {
		hour += 12
	}

	if period == "AM" && hour == 12 {
		hour = 0
	}

	// =======================
	// APPEND STRING TO SLICE
	// =======================
	newTimeParts := []string{} // slice kosong
	newTimeParts = append(newTimeParts, fmt.Sprintf("%02d", hour))
	newTimeParts = append(newTimeParts, minuteStr)
	newTimeParts = append(newTimeParts, secondStr)
	fmt.Println("Appended parts:", newTimeParts)

	// Gabungkan jadi format 24 jam
	converted := strings.Join(newTimeParts, ":")
	fmt.Println("Final 24-hour format:", converted)
}
