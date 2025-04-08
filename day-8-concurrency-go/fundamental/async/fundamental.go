// concurrency_fundamental.go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Fungsi sederhana untuk simulasi pekerjaan
func worker(id int, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()         // Kurangi counter WaitGroup setelah selesai
	time.Sleep(time.Second) // Simulasi delay
	result := fmt.Sprintf("Worker %d selesai bekerja", id)
	ch <- result // Kirim hasil ke channel
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 5) // Channel buffered agar tidak blocking

	// Menjalankan 5 worker secara concurrent
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg, ch)
	}

	// Tunggu semua goroutine selesai
	wg.Wait()
	close(ch) // Tutup channel setelah semua worker selesai

	// Ambil hasil dari channel
	for msg := range ch {
		fmt.Println(msg)
	}

	fmt.Println("Semua pekerjaan selesai!")
}
