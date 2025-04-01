package main

import "fmt"

// Node struct untuk linked list
type Node struct {
	Value int
	Next  *Node
}

func main() {
	// Membuat linked list: 1 -> 2 -> 3 -> nil
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}

	n1.Next = n2 // Node 1 menunjuk ke Node 2
	n2.Next = n3 // Node 2 menunjuk ke Node 3

	// Traversing linked list
	n := n1
	fmt.Print("Linked List:")
	for n != nil {
		fmt.Print(" ", n.Value)
		n = n.Next
	}
	fmt.Println()
}
