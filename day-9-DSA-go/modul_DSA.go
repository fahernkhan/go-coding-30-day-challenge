package main

import (
	"container/heap"
	"fmt"
)

// 1. Linked List
type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
}

func (ll *LinkedList) Insert(value int) {
	newNode := &LinkedListNode{Value: value}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (ll *LinkedList) Delete(value int) {
	if ll.Head == nil {
		return
	}
	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		return
	}
	current := ll.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

// 2. Stack
type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// 3. Queue
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// 4. Binary Search Tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Insert(value int) *TreeNode {
	if t == nil {
		return &TreeNode{Value: value}
	}
	if value < t.Value {
		t.Left = t.Left.Insert(value)
	} else {
		t.Right = t.Right.Insert(value)
	}
	return t
}

func (t *TreeNode) Search(value int) bool {
	if t == nil {
		return false
	}
	if t.Value == value {
		return true
	}
	if value < t.Value {
		return t.Left.Search(value)
	}
	return t.Right.Search(value)
}

func InOrder(root *TreeNode) {
	if root != nil {
		InOrder(root.Left)
		fmt.Printf("%d ", root.Value)
		InOrder(root.Right)
	}
}

// 5. Graph (Adjacency List)
type Graph struct {
	Vertices int
	AdjList  map[int][]int
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		AdjList:  make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.AdjList[u] = append(g.AdjList[u], v)
	g.AdjList[v] = append(g.AdjList[v], u) // Undirected graph
}

// 6. Heap (Priority Queue)
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 7. Hash Table
type HashTable struct {
	size  int
	table []*LinkedListNode
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:  size,
		table: make([]*LinkedListNode, size),
	}
}

func (ht *HashTable) hash(key int) int {
	return key % ht.size
}

func (ht *HashTable) Put(key, value int) {
	index := ht.hash(key)
	node := &LinkedListNode{Value: value}
	if ht.table[index] == nil {
		ht.table[index] = node
	} else {
		current := ht.table[index]
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

// 8. Algorithms
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// 9. Dynamic Programming (Fibonacci with Memoization)
var memo = make(map[int]int)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = Fibonacci(n-1) + Fibonacci(n-2)
	return memo[n]
}

func main() {
	// Contoh penggunaan semua modul
	fmt.Println("=== Linked List ===")
	ll := LinkedList{}
	ll.Insert(5)
	ll.Insert(10)
	ll.Insert(15)
	ll.Print()
	ll.Delete(10)
	ll.Print()

	fmt.Println("\n=== Stack ===")
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop())

	fmt.Println("\n=== Queue ===")
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	fmt.Println("Dequeue:", queue.Dequeue())
	fmt.Println("Dequeue:", queue.Dequeue())

	fmt.Println("\n=== BST ===")
	var root *TreeNode
	root = root.Insert(5)
	root.Insert(3)
	root.Insert(7)
	fmt.Print("InOrder: ")
	InOrder(root)
	fmt.Println("\nSearch 7:", root.Search(7))

	fmt.Println("\n=== Graph ===")
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	fmt.Println("Adjacency List:", g.AdjList)

	fmt.Println("\n=== Heap ===")
	h := &MinHeap{5, 2, 7}
	heap.Init(h)
	heap.Push(h, 1)
	fmt.Println("Heap Pop:", heap.Pop(h))

	fmt.Println("\n=== Hash Table ===")
	ht := NewHashTable(10)
	ht.Put(1, 100)
	ht.Put(11, 200)
	fmt.Println("Hash Table Bucket 1:", ht.table[1].Value)

	fmt.Println("\n=== Algorithms ===")
	arr := []int{5, 3, 8, 1}
	BubbleSort(arr)
	fmt.Println("Sorted Array:", arr)
	fmt.Println("Binary Search 8:", BinarySearch(arr, 8))

	fmt.Println("\n=== Dynamic Programming ===")
	fmt.Println("Fibonacci(10):", Fibonacci(10))
}
