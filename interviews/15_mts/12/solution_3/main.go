package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func NewNode(data int) *Node {

	node := new(Node)

	node.Data = data
	node.Left = nil
	node.Right = nil

	return node
}

func (n *Node) Traversal(result *[]int) {
	if n == nil {
		return
	}
	*result = append(*result, n.Data)
	n.Left.Traversal(result)
	n.Right.Traversal(result)
}
	

// Написать функцию для обхода дерева в глубину
func (n *Node) DFS() {
	var result []int
	n.Traversal(&result)
	fmt.Println(result)
}

func main() {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right.Left = NewNode(6)
	root.Right.Right = NewNode(7)

	root.DFS()
}
