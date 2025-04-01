package main

import "fmt"

func main() {
	// Array: Menyimpan jumlah elemen tetap dengan tipe yang sama
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println("Array:", a) // Output: [1 2 3]

	// Slice: Menyimpan jumlah elemen variabel dengan tipe yang sama
	s := []int{1, 2, 3}
	s = append(s, 4)
	fmt.Println("Slice:", s) // Output: [1 2 3 4]

	// Slice menggunakan make
	s2 := make([]int, 3) // Membuat slice dengan panjang 3
	s2[0] = 10
	s2[1] = 20
	s2[2] = 30
	// error jika dijalankan
	// s2[4] = 40
	fmt.Println("Slice dengan make:", s2) // Output: [10 20 30]

	s3 := make([]int, 6) // Panjang 0, kapasitas default
	s3 = append(s3, 1, 2, 3, 4, 5)
	fmt.Println("slice kapasitas default ditambahkan:", s3) // Output: [1 2 3 4 5]

	// Map: Menyimpan pasangan key-value
	m := map[string]int{"foo": 1, "bar": 2}
	m["baz"] = 3
	fmt.Println("Map:", m) // Output: map[bar:2 baz:3 foo:1]

	// Map menggunakan make
	m2 := make(map[string]int)
	m2["x"] = 100
	m2["y"] = 200
	fmt.Println("Map dengan make:", m2) // Output: map[x:100 y:200]

	// Struct: Menyimpan koleksi field terkait
	type User struct {
		ID   int
		Name string
	}
	u := User{ID: 1, Name: "Alice"}
	fmt.Println("Struct:", u) // Output: {1 Alice}

	// Pointer ke struct
	up := &u
	up.Name = "Bob"
	fmt.Println("Pointer ke struct:", *up) // Output: {1 Bob}

	// Channel: Digunakan untuk komunikasi antar goroutine
	ch := make(chan int, 2) // Channel dengan buffer 2
	ch <- 10
	ch <- 20
	fmt.Println("Channel:", <-ch, <-ch) // Output: 10 20

	// Linked List sederhana menggunakan struct
	type Node struct {
		Value int
		Next  *Node
	}
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n1.Next = n2
	n2.Next = n3

	// Traversing Linked List
	n := n1 // Menyimpan referensi ke node pertama (head)
	fmt.Print("Linked List:")
	for n != nil { // Selama n tidak nil (masih ada node)
		fmt.Print(" ", n.Value) // Cetak nilai node saat ini
		n = n.Next              // Pindah ke node berikutnya
	}
	fmt.Println()
}
