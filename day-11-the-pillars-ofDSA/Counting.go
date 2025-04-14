// 4. Counting – Frekuensi Kemunculan Kata
package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go go golang is awesome and go is fun"
	words := strings.Split(text, " ")

	count := make(map[string]int)
	for _, word := range words {
		count[word]++
	}

	fmt.Println("Frekuensi kata:")
	for word, freq := range count {
		fmt.Printf("%s: %d\n", word, freq)
	}
}
