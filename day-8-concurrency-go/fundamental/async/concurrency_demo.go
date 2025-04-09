package main

import (
	"fmt"
	"sync"
	"time"
)

// Fungsi untuk dijalankan sebagai goroutine
func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		fmt.Println(msg, "-", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	// Menjalankan beberapa goroutine
	wg.Add(2)
	go printMessage("Halo dari Goroutine 1", &wg)
	go printMessage("Halo dari Goroutine 2", &wg)

	// Channel untuk komunikasi antar goroutine
	ch := make(chan string)

	// Goroutine untuk mengirim data
	go func() {
		ch <- "Data dari channel 1"
	}()

	// Menerima data dari channel
	data := <-ch
	fmt.Println("Menerima:", data)

	// Buffered channel
	bufCh := make(chan string, 2)
	bufCh <- "Buffered 1"
	bufCh <- "Buffered 2"

	fmt.Println("Buffered:", <-bufCh)
	fmt.Println("Buffered:", <-bufCh)

	// Menggunakan select untuk multiple channel
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Data dari ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Data dari ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Menerima dari ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Menerima dari ch2:", msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout!")
	}

	wg.Wait()
	fmt.Println("Selesai semua goroutine.")
}
