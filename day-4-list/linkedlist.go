package main

import "fmt"

// Node class
type Node struct {
	property int
	nextNode *Node
}

// LinkedList class
type LinkedList struct {
	headNode *Node
}

// AddToHead method of LinkedList class
func (linkedList *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}

	linkedList.headNode = node
}

// ItarateList method iterates over LinkedList
func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

// Last method returns the last Node
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

// AddToEnd method adds the node with property to the end
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node
	lastNode = linkedList.LastNode()
	// misal last node dah ke-isi
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

// NodeWithValue method returns Node given parameter property
func (linkedlist LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedlist.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter method adds a node with nodeProperty after node with property
func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	// Mencari node dengan nilai nodeProperty
	var nodeWith *Node
	nodeWith = linkedList.NodeWithValue(nodeProperty)

	// Jika node ditemukan
	if nodeWith != nil {
		// Langkah 1: Hubungkan node baru ke node setelah nodeWith
		node.nextNode = nodeWith.nextNode
		// Langkah 2: Hubungkan nodeWith ke node baru
		nodeWith.nextNode = node
	}

}

func main() {
	var linkedList LinkedList

	linkedList = LinkedList{}

	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToHead(6)
	// fmt.Println(linkedList.headNode.property)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	linkedList.AddAfter(7, 24)
	linkedList.AddToEnd(100)

	// Expected list is [6, 3, 1, 7, 24, 5, 100]
	linkedList.IterateList()

}
