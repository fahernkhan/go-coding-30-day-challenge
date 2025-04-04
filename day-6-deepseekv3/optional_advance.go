package main

import (
	"fmt"
	"sort"
)

// ========== Advanced Data Structures ==========

// 1. AVL Tree (Self-Balancing BST)
type AVLNode struct {
	key    int
	height int
	left   *AVLNode
	right  *AVLNode
}

func (n *AVLNode) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *AVLNode) updateHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

func (n *AVLNode) balanceFactor() int {
	return n.left.getHeight() - n.right.getHeight()
}

func rotateRight(y *AVLNode) *AVLNode {
	x := y.left
	T := x.right

	x.right = y
	y.left = T

	y.updateHeight()
	x.updateHeight()
	return x
}

func rotateLeft(x *AVLNode) *AVLNode {
	y := x.right
	T := y.left

	y.left = x
	x.right = T

	x.updateHeight()
	y.updateHeight()
	return y
}

func insertAVL(root *AVLNode, key int) *AVLNode {
	if root == nil {
		return &AVLNode{key: key, height: 1}
	}

	if key < root.key {
		root.left = insertAVL(root.left, key)
	} else if key > root.key {
		root.right = insertAVL(root.right, key)
	} else {
		return root
	}

	root.updateHeight()
	balance := root.balanceFactor()

	// Left Left Case
	if balance > 1 && key < root.left.key {
		return rotateRight(root)
	}
	// Right Right Case
	if balance < -1 && key > root.right.key {
		return rotateLeft(root)
	}
	// Left Right Case
	if balance > 1 && key > root.left.key {
		root.left = rotateLeft(root.left)
		return rotateRight(root)
	}
	// Right Left Case
	if balance < -1 && key < root.right.key {
		root.right = rotateRight(root.right)
		return rotateLeft(root)
	}

	return root
}

// 2. Trie (Prefix Tree)
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func NewTrie() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func (t *TrieNode) Insert(word string) {
	node := t
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = NewTrie()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *TrieNode) Search(word string) bool {
	node := t
	for _, ch := range word {
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

// 3. Union-Find (Disjoint Set Union)
type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(size int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, size),
		rank:   make([]int, size),
	}
	for i := range uf.parent {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else {
			uf.parent[rootX] = rootY
			if uf.rank[rootX] == uf.rank[rootY] {
				uf.rank[rootY]++
			}
		}
	}
}

// ========== Advanced Algorithms ==========

// 1. Kruskal's Algorithm (MST)
type Edge struct {
	u, v, weight int
}

func kruskalMST(edges []Edge, vertices int) []Edge {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	uf := NewUnionFind(vertices)
	var result []Edge

	for _, edge := range edges {
		if uf.Find(edge.u) != uf.Find(edge.v) {
			result = append(result, edge)
			uf.Union(edge.u, edge.v)
		}
	}
	return result
}

// 2. Topological Sort (Kahn's Algorithm)
func topologicalSort(graph map[int][]int, inDegree map[int]int) []int {
	queue := []int{}
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	var result []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return result
}

// 3. KMP Algorithm (String Matching)
func computeLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0

	for i := 1; i < len(pattern); {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func KMP(text, pattern string) []int {
	lps := computeLPS(pattern)
	var matches []int

	i, j := 0, 0
	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++
		}

		if j == len(pattern) {
			matches = append(matches, i-j)
			j = lps[j-1]
		} else if i < len(text) && text[i] != pattern[j] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return matches
}

// 4. Segment Tree (Range Sum Query)
type SegmentTree struct {
	n    int
	tree []int
}

func NewSegmentTree(arr []int) *SegmentTree {
	st := &SegmentTree{
		n:    len(arr),
		tree: make([]int, 2*len(arr)),
	}
	for i := 0; i < len(arr); i++ {
		st.tree[len(arr)+i] = arr[i]
	}
	for i := len(arr) - 1; i > 0; i-- {
		st.tree[i] = st.tree[2*i] + st.tree[2*i+1]
	}
	return st
}

func (st *SegmentTree) Update(pos, value int) {
	pos += st.n
	st.tree[pos] = value
	for pos > 1 {
		pos >>= 1
		st.tree[pos] = st.tree[2*pos] + st.tree[2*pos+1]
	}
}

func (st *SegmentTree) Query(l, r int) int {
	res := 0
	l += st.n
	r += st.n

	for l <= r {
		if l%2 == 1 {
			res += st.tree[l]
			l++
		}
		if r%2 == 0 {
			res += st.tree[r]
			r--
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func main() {
	// AVL Tree Example
	var avlRoot *AVLNode
	avlRoot = insertAVL(avlRoot, 10)
	avlRoot = insertAVL(avlRoot, 20)
	avlRoot = insertAVL(avlRoot, 30)
	fmt.Println("AVL Tree Root:", avlRoot.key)

	// Trie Example
	trie := NewTrie()
	trie.Insert("apple")
	trie.Insert("app")
	fmt.Println("Trie Search 'app':", trie.Search("app"))

	// Kruskal's MST Example
	edges := []Edge{
		{0, 1, 10},
		{0, 2, 6},
		{0, 3, 5},
		{1, 3, 15},
		{2, 3, 4},
	}
	fmt.Println("Kruskal's MST Edges:", kruskalMST(edges, 4))

	// Topological Sort Example
	graph := map[int][]int{
		5: {2, 0},
		4: {0, 1},
		2: {3},
		3: {1},
	}
	inDegree := map[int]int{0: 2, 1: 2, 2: 1, 3: 1, 5: 0, 4: 0}
	fmt.Println("Topological Order:", topologicalSort(graph, inDegree))

	// KMP Example
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"
	fmt.Println("KMP Matches at indices:", KMP(text, pattern))

	// Segment Tree Example
	arr := []int{1, 3, 5, 7, 9, 11}
	st := NewSegmentTree(arr)
	fmt.Println("Segment Tree Query [1-3]:", st.Query(1, 3))
	st.Update(1, 10)
	fmt.Println("After Update, Query [1-3]:", st.Query(1, 3))
}
