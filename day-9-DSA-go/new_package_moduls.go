package main

import (
	"container/heap"
	"fmt"
	"maps"
	"math"
	"math/rand"
	"slices"
	"sort"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 1. Package slices - Manipulasi Slice
	numbers := []int{5, 2, 8, 1, 3}
	fmt.Println("Original slice:", numbers)

	slices.Sort(numbers)
	fmt.Println("Sorted slice:", numbers)

	_, found := slices.BinarySearch(numbers, 8)
	fmt.Println("Binary search for 8:", found)

	slices.Reverse(numbers)
	fmt.Println("Reversed slice:", numbers)

	// 2. Package maps - Operasi Map
	userMap := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 28}
	fmt.Println("\nMap keys:", maps.Keys(userMap))
	fmt.Println("Map values:", maps.Values(userMap))

	// 3. Package math & cmp - Operasi Matematika
	fmt.Println("\nMath operations:")
	fmt.Println("Max between 7 and 3:", max(7, 3))
	fmt.Println("Min between 4.5 and 5.2:", min(4.5, 5.2))
	fmt.Println("Absolute value of -15:", math.Abs(-15))
	fmt.Println("Square root of 25:", math.Sqrt(25))

	// 4. Priority Queue dengan container/heap
	fmt.Println("\nPriority Queue:")
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{value: "Orange", priority: 3})
	heap.Push(pq, &Item{value: "Apple", priority: 1})
	heap.Push(pq, &Item{value: "Banana", priority: 2})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		fmt.Printf("%s (Priority: %d)\n", item.value, item.priority)
	}

	// 5. Custom Sorting dengan sort
	fmt.Println("\nCustom sorting:")
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted by age:", people)

	// 6. String Manipulation dengan strings
	fmt.Println("\nString operations:")
	s := "Hello, World!"
	fmt.Println("Uppercase:", strings.ToUpper(s))
	fmt.Println("Split:", strings.Split(s, " "))
	fmt.Println("Contains 'World':", strings.Contains(s, "World"))

	// 7. Algoritma BFS dengan Slice
	fmt.Println("\nBFS Algorithm:")
	graph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {4},
		3: {},
		4: {5},
		5: {},
	}
	bfs(graph, 0)

	// 8. Random Number Generation
	fmt.Println("\nRandom numbers:")
	fmt.Println("Random int:", rand.Intn(100))
	fmt.Println("Shuffled slice:", rand.Perm(10))
}

// Implementasi Priority Queue
type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Implementasi BFS
func bfs(graph map[int][]int, start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		if !visited[vertex] {
			fmt.Printf("Visited %d\n", vertex)
			visited[vertex] = true
			queue = append(queue, graph[vertex]...)
		}
	}
}
