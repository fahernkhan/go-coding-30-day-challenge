package main

import (
	"container/list"
	"fmt"
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
	// fmt.Println()
}
