package main

import (
	"between-two-sets/internal/logic" // Ganti dengan nama module kamu
	"fmt"
)

func main() {
	a := []int32{2, 4}
	b := []int32{16, 32, 96}

	result := logic.GetTotalX(a, b)
	fmt.Println("Jumlah x yang valid:", result)
}
