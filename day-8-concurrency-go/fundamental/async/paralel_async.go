package main

import (
	"fmt"
	"time"
)

func printNumber(n int) {
	fmt.Printf("Printing number: %d\n", n)
}

func main() {
	for i := 1; i <= 5; i++ {
		go printNumber(i) // Setiap iterasi membuat goroutine baru
	}

	time.Sleep(1 * time.Second) // Beri waktu semua goroutine selesai
	fmt.Println("All goroutines launched.")
}
