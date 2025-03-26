package main

import "fmt"

// Node struct merepresentasikan satu node dalam LinkedList
type Node struct {
	property int   // Nilai yang disimpan di dalam node
	nextNode *Node // Pointer ke node berikutnya
}

// LinkedListInterface adalah interface untuk operasi dasar LinkedList
type LinkedListInterface interface {
	AddToHead(property int)                  // Menambahkan node di awal
	AddToEnd(property int)                   // Menambahkan node di akhir
	IterateList()                            // Menampilkan semua node
	AddAfter(nodeProperty int, property int) // Menambahkan node setelah node tertentu
}

// LinkedList struct yang mengimplementasikan LinkedListInterface
type LinkedList struct {
	headNode *Node
}

// AddToHead menambahkan node baru di awal LinkedList
func (linkedList *LinkedList) AddToHead(property int) {
	// Membuat node baru
	var node = &Node{property: property, nextNode: linkedList.headNode}
	linkedList.headNode = node
}

// IterateList menampilkan seluruh elemen dalam LinkedList
func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

// LastNode mengembalikan node terakhir dalam LinkedList
func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd menambahkan node baru di akhir LinkedList
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{property: property, nextNode: nil}
	var lastNode *Node = linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	} else {
		linkedList.headNode = node // Jika LinkedList kosong
	}
}

// NodeWithValue mencari node berdasarkan property tertentu
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			return node
		}
	}
	return nil
}

// AddAfter menambahkan node baru setelah node dengan nilai tertentu
func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var nodeWith *Node = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		var node = &Node{property: property, nextNode: nodeWith.nextNode}
		nodeWith.nextNode = node
	}
}

func main() {
	// Menggunakan LinkedList sebagai LinkedListInterface
	var linkedList LinkedListInterface = &LinkedList{}

	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToHead(6)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	linkedList.AddAfter(7, 24)
	linkedList.AddToEnd(100)

	// Expected list: [6, 3, 1, 7, 24, 5, 100]
	linkedList.IterateList()
}
