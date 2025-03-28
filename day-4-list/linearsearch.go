package main

import "fmt"

func LinearSearch(elements []int, findElement int) bool {
	var element int
	for _, element = range elements {
		if element == findElement {
			return true
		}
	}
	return false
}

func main() {
	var elements []int
	elements = []int{15, 48, 26, 18, 41, 86, 29, 51, 20}
	fmt.Println(LinearSearch(elements, 24))
	fmt.Println(LinearSearch(elements, 51))

}
