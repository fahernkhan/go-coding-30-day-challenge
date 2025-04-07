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
