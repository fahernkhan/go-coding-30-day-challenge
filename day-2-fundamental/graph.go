package main

import "fmt"

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DFS(visit func(n *NOde)) {
	visit(n)
	for _, child := range n.Children {
		child.DFS(visit)
	}
}

func main() {
	root := &Node{
		Name: "root",
		Children: []*Node{
			{Name: "child1"},
			{Name: "child2"},
			{Name: "child3"},
		},
	}

	// Perform a depth-first search of the graph, printing the name of each node
	root.DFS(func(n *Node) {
		fmt.Println(n.Name)
	})
	// Output: root, child1, child2, child3
}
