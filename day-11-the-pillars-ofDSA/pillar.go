package main

import (
	"container/heap"
	"fmt"
)

// ===============================
// 1. ARRAY / SLICE
// ===============================
func demoArray() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Array/Slice:", arr)
}

// ===============================
// 2. HASHMAP / MAP
// ===============================
func demoMap() {
	m := map[string]int{
		"apple":  5,
		"banana": 3,
	}
	fmt.Println("HashMap:", m)
}

// ===============================
// 3. QUEUE (FIFO)
// ===============================
type Queue []int

func (q *Queue) Enqueue(x int) {
	*q = append(*q, x)
}

func (q *Queue) Dequeue() int {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func demoQueue() {
	var q Queue
	q.Enqueue(10)
	q.Enqueue(20)
	fmt.Println("Queue Dequeue:", q.Dequeue())
}

// ===============================
// 4. STACK (LIFO)
// ===============================
type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	n := len(*s)
	val := (*s)[n-1]
	*s = (*s)[:n-1]
	return val
}

func demoStack() {
	var s Stack
	s.Push(100)
	s.Push(200)
	fmt.Println("Stack Pop:", s.Pop())
}

// ===============================
// 5. HEAP / PRIORITY QUEUE
// ===============================
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] } // Min-Heap
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[0 : n-1]
	return val
}

func demoHeap() {
	h := &IntHeap{5, 1, 3}
	heap.Init(h)
	heap.Push(h, 2)
	fmt.Println("Heap Pop:", heap.Pop(h))
}

// ===============================
// 6. BINARY SEARCH
// ===============================
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func demoBinarySearch() {
	arr := []int{1, 3, 5, 7, 9}
	fmt.Println("Binary Search index:", binarySearch(arr, 5))
}

// ===============================
// 7. GRAPH (Adjacency List + DFS)
// ===============================
type Graph map[int][]int

func dfs(g Graph, start int, visited map[int]bool) {
	if visited[start] {
		return
	}
	visited[start] = true
	fmt.Print(start, " ")

	for _, neighbor := range g[start] {
		dfs(g, neighbor, visited)
	}
}

func demoGraph() {
	g := Graph{
		0: {1, 2},
		1: {0, 3},
		2: {0},
		3: {1},
	}
	visited := make(map[int]bool)
	fmt.Print("Graph DFS: ")
	dfs(g, 0, visited)
	fmt.Println()
}

// ===============================
// MAIN FUNCTION
// ===============================
func main() {
	fmt.Println("===== DSA/ASD Fundamentals Demo in GoLang =====")
	demoArray()
	demoMap()
	demoQueue()
	demoStack()
	demoHeap()
	demoBinarySearch()
	demoGraph()
}
