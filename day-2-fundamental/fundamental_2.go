package main

import (
	"container/list"
	"fmt"
)

// Linear Search
func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

// Binary Search (assumes sorted array)
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
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

// Bubble Sort
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// Stack Implementation using List
type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{list: list.New()}
}

func (s *Stack) Push(value any) {
	s.list.PushBack(value)
}

func (s *Stack) Pop() any {
	e := s.list.Back()
	if e != nil {
		s.list.Remove(e)
		return e.Value
	}
	return nil
}

// Queue Implementation using Slice
type Queue struct {
	elements []any
}

func (q *Queue) Enqueue(value any) {
	q.elements = append(q.elements, value)
}

func (q *Queue) Dequeue() any {
	if len(q.elements) == 0 {
		return nil
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value
}

// Graph using Adjacency List
type Graph struct {
	adjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node, " ")
		for _, neighbor := range g.adjList[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 2}
	bubbleSort(arr)
	fmt.Println("Sorted Array:", arr)

	fmt.Println("Binary Search 5:", binarySearch(arr, 5))
	fmt.Println("Linear Search 4:", linearSearch(arr, 4))

	stack := NewStack()
	stack.Push(10)
	stack.Push(20)
	fmt.Println("Stack Pop:", stack.Pop())

	queue := Queue{}
	queue.Enqueue(30)
	queue.Enqueue(40)
	fmt.Println("Queue Dequeue:", queue.Dequeue())

	graph := NewGraph()
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 5)
	fmt.Println("BFS Traversal:")
	graph.BFS(1)
}
