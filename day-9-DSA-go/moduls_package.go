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
)

func main() {
	// ==============================================
	// 1. Paket "slices" (Go 1.21+)
	// ==============================================
	nums := []int{5, 2, 8, 1, 3}

	// Sorting
	slices.Sort(nums)
	fmt.Println("Sorted slice:", nums) // [1 2 3 5 8]

	// Binary Search
	idx, found := slices.BinarySearch(nums, 5)
	fmt.Printf("Binary search: index %d, found %t\n", idx, found) // index 3, true

	// Reverse
	slices.Reverse(nums)
	fmt.Println("Reversed:", nums) // [8 5 3 2 1]

	// ==============================================
	// 2. Paket "math"
	// ==============================================
	fmt.Println("\nMath operations:")
	fmt.Println("Max of 3 and 5:", max(3, 5))     // 5
	fmt.Println("Min of 3 and 5:", min(3, 5))     // 3
	fmt.Println("Abs of -4.5:", math.Abs(-4.5))   // 4.5
	fmt.Println("Ceil of 3.2:", math.Ceil(3.2))   // 4
	fmt.Println("Floor of 3.8:", math.Floor(3.8)) // 3
	fmt.Println("Pow 2^3:", math.Pow(2, 3))       // 8

	// ==============================================
	// 3. Paket "container/heap" (Priority Queue)
	// ==============================================
	fmt.Println("\nHeap/Priority Queue:")
	// Min-heap example
	h := &IntHeap{3, 1, 4, 1, 5}
	heap.Init(h)
	heap.Push(h, 2)
	fmt.Println("Heap elements:")
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h)) // 1 1 2 3 4 5
	}

	// ==============================================
	// 4. Paket "sort"
	// ==============================================
	fmt.Println("\n\nCustom sorting:")
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	// Sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted by age:", people)

	// ==============================================
	// 5. Paket "strings" dan "bytes"
	// ==============================================
	fmt.Println("\nString operations:")
	s := "Hello, DSA!"
	fmt.Println("Contains 'DSA':", strings.Contains(s, "DSA")) // true
	fmt.Println("To uppercase:", strings.ToUpper(s))           // HELLO, DSA!
	fmt.Println("Split:", strings.Split(s, " "))               // [Hello, DSA!]

	// ==============================================
	// 6. Paket "maps" (Go 1.21+)
	// ==============================================
	fmt.Println("\nMap operations:")
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println("Map keys:", maps.Keys(m))     // [a b] (unordered)
	fmt.Println("Map values:", maps.Values(m)) // [1 2] (unordered)

	// ==============================================
	// 7. Paket "rand"
	// ==============================================
	fmt.Println("\nRandom numbers:")
	rand.Seed(42)                              // Seed untuk reproducibility
	fmt.Println("Random int:", rand.Intn(100)) // Angka acak 0-99
	fmt.Println("Random perm:", rand.Perm(5))  // Permutasi acak [0 1 2 3 4]

	// Contoh pencarian maksimum dalam slice
	maxVal := slices.Max(nums)

	// Contoh shuffle data
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	// Contoh penggunaan math untuk DSA
	distance := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// ==============================================
// Implementasi untuk heap
// ==============================================
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// ==============================================
// Struct untuk contoh sorting
// ==============================================
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}
