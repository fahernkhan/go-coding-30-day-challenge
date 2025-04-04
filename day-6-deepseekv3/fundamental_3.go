package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
)

// ========== Basic Data Structures ==========

// 1. Array & Slice
func arraySliceExample() {
	// Fixed size array
	arr := [3]int{1, 2, 3}

	// Dynamic slice
	slice := []int{1, 2, 3}
	slice = append(slice, 4)

	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)
}

// 2. Linked List (Singly)
type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Append(data int) {
	newNode := &Node{Data: data}

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

// 3. Stack (LIFO)
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

// 4. Queue (FIFO)
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

// ========== Basic Algorithms ==========

// 1. Quick Sort
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var less, greater, equal []int

	for _, num := range arr {
		switch {
		case num < pivot:
			less = append(less, num)
		case num > pivot:
			greater = append(greater, num)
		default:
			equal = append(equal, num)
		}
	}

	return append(append(quickSort(less), equal...), quickSort(greater)...)
}

// 2. Binary Search
func binarySearch(arr []int, target int) int {
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

// 3. BFS (Breadth First Search)
func bfs(graph map[int][]int, start int) {
	visited := make(map[int]bool)
	queue := list.New()
	queue.PushBack(start)

	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value.(int)
		queue.Remove(element)

		if !visited[node] {
			visited[node] = true
			fmt.Printf("%d ", node)

			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					queue.PushBack(neighbor)
				}
			}
		}
	}
}

// ========== Advanced Data Structures ==========
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = insertBST(root.Left, val)
	} else {
		root.Right = insertBST(root.Right, val)
	}
	return root
}

func searchBST(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}

	if root.Val == val {
		return true
	} else if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

// 2. Graph (DFS & Dijkstra's Algorithm)
func dfs(graph map[int][]int, node int, visited map[int]bool) {
	if visited[node] {
		return
	}

	visited[node] = true
	fmt.Printf("%d ", node)

	for _, neighbor := range graph[node] {
		dfs(graph, neighbor, visited)
	}
}

// Dijkstra's Priority Queue implementation
type Edge struct {
	node   int
	weight int
}

type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].weight < pq[j].weight }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Edge))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func dijkstra(graph map[int][]Edge, start int) map[int]int {
	dist := make(map[int]int)
	for node := range graph {
		dist[node] = math.MaxInt32
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Edge{start, 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Edge)
		u := current.node

		for _, edge := range graph[u] {
			if dist[edge.node] > dist[u]+edge.weight {
				dist[edge.node] = dist[u] + edge.weight
				heap.Push(pq, &Edge{edge.node, dist[edge.node]})
			}
		}
	}
	return dist
}

// 3. Hash Table with Chaining
type HashTable struct {
	size    int
	buckets [][]keyValue
}

type keyValue struct {
	key   string
	value interface{}
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:    size,
		buckets: make([][]keyValue, size),
	}
}

func (ht *HashTable) hash(key string) int {
	sum := 0
	for _, c := range key {
		sum += int(c)
	}
	return sum % ht.size
}

func (ht *HashTable) Set(key string, value interface{}) {
	index := ht.hash(key)
	for i, kv := range ht.buckets[index] {
		if kv.key == key {
			ht.buckets[index][i].value = value
			return
		}
	}
	ht.buckets[index] = append(ht.buckets[index], keyValue{key, value})
}

func (ht *HashTable) Get(key string) interface{} {
	index := ht.hash(key)
	for _, kv := range ht.buckets[index] {
		if kv.key == key {
			return kv.value
		}
	}
	return nil
}

// 4. Min-Heap
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

// ========== Advanced Algorithms ==========
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// 2. Dynamic Programming (Coin Change)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	// Running examples
	arraySliceExample()

	// Linked List
	ll := LinkedList{}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	// Stack
	s := Stack{}
	s.Push(1)
	s.Push(2)
	fmt.Println("\nStack Pop:", s.Pop())

	// Quick Sort
	unsorted := []int{3, 1, 4, 1, 5, 9, 2, 6}
	sorted := quickSort(unsorted)
	fmt.Println("Quick Sort Result:", sorted)

	// Binary Search
	fmt.Println("Binary Search found at index:", binarySearch(sorted, 5))

	// BFS Example
	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
	}
	fmt.Print("BFS Traversal: ")
	bfs(graph, 1)
	fmt.Println()

	// ADVANCE
	// BST Example
	root := insertBST(nil, 5)
	insertBST(root, 3)
	insertBST(root, 7)
	fmt.Println("\nBST Search 7:", searchBST(root, 7))

	// DFS Example
	graph = map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
	}
	fmt.Print("DFS Traversal: ")
	dfs(graph, 1, make(map[int]bool))

	// Dijkstra's Example
	weightedGraph := map[int][]Edge{
		0: {{1, 4}, {2, 1}},
		1: {{3, 1}},
		2: {{1, 2}, {3, 5}},
		3: {},
	}
	fmt.Println("\nDijkstra Shortest Paths:", dijkstra(weightedGraph, 0))

	// HashTable Example
	ht := NewHashTable(10)
	ht.Set("foo", 42)
	ht.Set("bar", "hello")
	fmt.Println("HashTable Get 'foo':", ht.Get("foo"))

	// Merge Sort Example
	unsortedArr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Merge Sort Result:", mergeSort(unsortedArr))

	// Coin Change Example
	coins := []int{1, 2, 5}
	fmt.Println("Coin Change Min Coins:", coinChange(coins, 11))
}
