# 🧵 Fundamental Concurrency di Golang

## 🌀 Bagian 1: Goroutine — Dasar Eksekusi Asinkron

**Goroutine** adalah *unit eksekusi ringan* dalam bahasa Go yang memungkinkan kita menjalankan fungsi secara **asinkron** dan **concurrent** hanya dengan keyword `go`.

---

## 🧠 Konsep Dasar Goroutine

| Konsep         | Penjelasan                                                                 |
|----------------|------------------------------------------------------------------------------|
| `go` keyword   | Menjalankan fungsi sebagai goroutine                                         |
| Asynchronous   | Fungsi berjalan bersamaan tanpa blocking                                     |
| Non-blocking   | `main()` tidak menunggu goroutine selesai                                    |
| `time.Sleep`   | Digunakan untuk memberi waktu eksekusi goroutine                             |
| Race Condition | Bisa terjadi jika banyak goroutine mengakses data yang sama secara bersamaan |

---

## 💡 Contoh 1: Menjalankan Fungsi Biasa sebagai Goroutine

```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from Goroutine")
}

func main() {
    go sayHello() // Menjalankan sayHello sebagai goroutine
    fmt.Println("Main function continues...")
    time.Sleep(1 * time.Second) // Memberi waktu goroutine untuk selesai
}
```

**Penjelasan:**

- Fungsi `sayHello()` akan berjalan secara asinkron.
- Tanpa `time.Sleep`, program mungkin selesai sebelum goroutine dieksekusi.
- Ketika `main()` selesai, semua goroutine yang belum selesai akan dihentikan paksa.

---

## 💡 Contoh 2: Banyak Goroutine dalam Loop

```go
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
        go printNumber(i)
    }
    time.Sleep(1 * time.Second)
    fmt.Println("All goroutines launched.")
}
```

**Penjelasan:**

- Fungsi `printNumber(i)` dipanggil secara asinkron sebanyak 5 kali.
- Karena goroutine berjalan bersamaan, urutan output bisa acak.
- Gunakan `time.Sleep()` hanya sebagai simulasi sederhana; nanti akan belajar `sync.WaitGroup`.

---

## 💡 Contoh 3: Fungsi Anonim dalam Goroutine

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    go func(msg string) {
        fmt.Println("Message:", msg)
    }("Hello with anonymous function")

    time.Sleep(1 * time.Second)
}
```

**Penjelasan:**

- Fungsi anonim digunakan langsung sebagai goroutine.
- Parameter dapat diteruskan langsung ke dalam fungsi anonim.
- Berguna untuk eksekusi cepat tanpa membuat fungsi terpisah.

---

## ⚠️ Contoh 4: Goroutine Tanpa `time.Sleep`

```go
package main

import "fmt"

func main() {
    go fmt.Println("This might not print")
}
```

**Penjelasan:**

- Output mungkin tidak muncul sama sekali karena `main()` langsung selesai.
- Ini menunjukkan bahwa *lifetime* goroutine bergantung pada `main()`.
- Perlu mekanisme sinkronisasi agar goroutine punya cukup waktu untuk berjalan.

---

## 🔧 Tips Praktis

| Tips  | Penjelasan |
|-------|------------|
| ✅    | Gunakan goroutine untuk fungsi yang tidak harus menunggu hasil langsung |
| ✅    | Tambahkan `time.Sleep()` hanya untuk eksperimen sederhana |
| ✅    | Hindari penggunaan goroutine tanpa mekanisme kontrol seperti `WaitGroup`, `channel`, atau `context` |
| ⚠️    | Jangan anggap goroutine selesai tepat waktu tanpa sinkronisasi |
| ⚠️    | Waspadai *race condition* saat banyak goroutine akses data bersama |

---

## ✅ Kesimpulan

> **Goroutine** adalah fitur powerful dari Go untuk menjalankan fungsi secara paralel tanpa overhead besar.  
> Sangat ringan dibanding thread biasa.  
> Mudah digunakan, cukup dengan `go namaFungsi()`.  
> Tetapi: **jangan lupa kontrol eksekusinya!**  
> `time.Sleep()` bukan solusi jangka panjang. Gunakan `sync.WaitGroup`, `channel`, atau `context`.