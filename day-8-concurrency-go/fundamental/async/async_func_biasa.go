package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from Goroutine")
}

func main() {
	go sayHello() // Goroutine pertama
	fmt.Println("Main function continues...")
	time.Sleep(1 * time.Second) // Tunggu agar goroutine selesai
}
