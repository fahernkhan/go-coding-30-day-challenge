package main

import (
	"fmt"
)

// Fungsi untuk mengonversi angka Romawi ke bilangan desimal
func romanToInt(s string) int {
	// TODO: Implementasikan solusi kamu di sini
	romanChar := []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
	romanValue := []int{1, 5, 10, 50, 100, 500, 1000}

	getValue := func(roman byte) int {
		for i, value := range romanChar {
			if roman == value {
				return romanValue[i]
			}
		}
		return 0
	}

	total := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		currentValue := getValue(s[i])
		if currentValue < prevValue {
			total -= currentValue
		} else {
			total += currentValue
		}
		prevValue = currentValue
	}
	return total
}

// Fungsi untuk menjalankan unit test otomatis
func runTests() {
	testCases := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXC", 1990},
		{"MMXXIV", 2024},
		{"XLII", 42},
		{"CDXLIV", 444},
	}

	for _, tc := range testCases {
		result := romanToInt(tc.input)
		if result == tc.expected {
			fmt.Printf("✅ PASS: %s -> %d\n", tc.input, result)
		} else {
			fmt.Printf("❌ FAIL: %s -> Expected %d, but got %d\n", tc.input, tc.expected, result)
		}
	}
}

func main() {
	fmt.Println("🏆 Running Roman Numeral Conversion Tests...")
	runTests()
}
