package main

import (
	"container/heap"
	"fmt"
)

// ===============================
// 1. ARRAY/SLICE (Dynamic Array)
// ===============================
func demoArray() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("\n1. Array/Slice:")
	fmt.Println("Initial:", arr)

	// Slice tricks
	arr = append(arr, 6)              // Append
	sub := arr[2:4]                   // Sub-slicing
	arr = append(arr[:3], arr[4:]...) // Remove element
	fmt.Println("After manipulations:", arr, "Sub-slice:", sub)
}

// ===============================
// 2. HASHMAP/MAP
// ===============================
func demoMap() {
	m := map[string]int{
		"apple":  5,
		"banana": 3,
	}
	fmt.Println("\n2. HashMap:")
	fmt.Println("Initial:", m)

	m["orange"] = 4    // Add
	delete(m, "apple") // Delete
	fmt.Println("Updated:", m)
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
	fmt.Println("\n3. Queue:")
	fmt.Println("Dequeued:", q.Dequeue(), "Remaining:", q)
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
	fmt.Println("\n4. Stack:")
	fmt.Println("Popped:", s.Pop(), "Remaining:", s)
}

// ===============================
// 5. HEAP/PRIORITY QUEUE
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
	fmt.Println("\n5. Heap:")
	fmt.Println("Popped:", heap.Pop(h), "Remaining:", h)
}

// ===============================
// 6. BINARY SEARCH
// ===============================
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		switch {
		case arr[mid] == target:
			return mid
		case arr[mid] < target:
			low = mid + 1
		default:
			high = mid - 1
		}
	}
	return -1
}

func demoBinarySearch() {
	arr := []int{1, 3, 5, 7, 9}
	fmt.Println("\n6. Binary Search:")
	fmt.Println("Index of 5:", binarySearch(arr, 5))
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
	fmt.Println("\n7. Graph DFS:")
	fmt.Print("Traversal: ")
	dfs(g, 0, visited)
	fmt.Println()
}

// ===============================
// 8. LINKED LIST (SINGLY)
// ===============================
type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList() *ListNode {
	head := &ListNode{Val: 1}
	current := head
	for i := 2; i <= 5; i++ {
		current.Next = &ListNode{Val: i}
		current = current.Next
	}
	return head
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}

func demoLinkedList() {
	head := createLinkedList()
	fmt.Println("\n8. Linked List:")
	fmt.Print("Elements: ")
	printLinkedList(head)
	fmt.Println()
}

// ===============================
// 9. BINARY TREE (TRAVERSALS)
// ===============================
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(root *TreeNode) {
	if root != nil {
		inOrder(root.Left)
		fmt.Print(root.Val, " ")
		inOrder(root.Right)
	}
}

func createBinaryTree() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	return root
}

func demoBinaryTree() {
	root := createBinaryTree()
	fmt.Println("\n9. Binary Tree:")
	fmt.Print("In-Order Traversal: ")
	inOrder(root)
	fmt.Println()
}

// ===============================
// 10. GRAPH BFS
// ===============================
func bfs(g Graph, start int) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node, " ")
		for _, neighbor := range g[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func demoBFS() {
	g := Graph{
		0: {1, 2},
		1: {0, 3},
		2: {0},
		3: {1},
	}
	fmt.Println("\n10. Graph BFS:")
	fmt.Print("Traversal: ")
	bfs(g, 0)
	fmt.Println()
}

// ===============================
// 11. QUICKSORT
// ===============================
func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[len(arr)/2]
	var less, equal, greater []int
	for _, n := range arr {
		switch {
		case n < pivot:
			less = append(less, n)
		case n == pivot:
			equal = append(equal, n)
		default:
			greater = append(greater, n)
		}
	}
	return append(append(quicksort(less), append(equal, quicksort(greater)...)...))
}

func demoQuicksort() {
	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	fmt.Println("\n11. Quicksort:")
	fmt.Println("Unsorted:", arr)
	fmt.Println("Sorted:  ", quicksort(arr))
}

// ===============================
// 12. DYNAMIC PROGRAMMING (FIBONACCI)
// ===============================
func fibonacci(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)
	return memo[n]
}

func demoDP() {
	memo := make(map[int]int)
	fmt.Println("\n12. Dynamic Programming:")
	fmt.Println("Fibonacci(10):", fibonacci(10, memo))
}

// ===============================
// 13. GO TRICKS
// ===============================
func demoGoTricks() {
	fmt.Println("\n13. Go Tricks:")

	// Named return values
	split := func(sum int) (x, y int) {
		x = sum * 4 / 9
		y = sum - x
		return
	}
	a, b := split(10)
	fmt.Println("Named Returns:", a, b)

	// Defer for LIFO execution
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer (executed first!)")

	// Slice reversal
	reverse := func(s []int) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}
	arr := []int{1, 2, 3, 4, 5}
	reverse(arr)
	fmt.Println("Reversed Slice:", arr)
}

// ===============================
// MAIN FUNCTION
// ===============================
func main() {
	fmt.Println("===== DSA Fundamentals in Go =====")
	demoArray()
	demoMap()
	demoQueue()
	demoStack()
	demoHeap()
	demoBinarySearch()
	demoGraph()
	demoLinkedList()
	demoBinaryTree()
	demoBFS()
	demoQuicksort()
	demoDP()
	demoGoTricks()

	fmt.Println("\n===== End of Demo =====")
}
