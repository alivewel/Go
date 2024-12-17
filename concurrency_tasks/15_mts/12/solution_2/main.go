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

// Прямой обход (pre-order)
func (n *Node) DFSPreOrder() {
	if n == nil {
		return
	}
	fmt.Print(n.Data, " ") // Обработка узла
	n.Left.DFSPreOrder()    // Обход левого поддерева
	n.Right.DFSPreOrder()   // Обход правого поддерева
}

// Симметричный обход (in-order)
func (n *Node) DFSInOrder() {
	if n == nil {
		return
	}
	n.Left.DFSInOrder()     // Обход левого поддерева
	fmt.Print(n.Data, " ")  // Обработка узла
	n.Right.DFSInOrder()    // Обход правого поддерева
}

// Обратный обход (post-order)
func (n *Node) DFSPostOrder() {
	if n == nil {
		return
	}
	n.Left.DFSPostOrder()   // Обход левого поддерева
	n.Right.DFSPostOrder()  // Обход правого поддерева
	fmt.Print(n.Data, " ")  // Обработка узла
}

func main() {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right.Left = NewNode(6)
	root.Right.Right = NewNode(7)

	fmt.Println("DFS Pre-Order:")
	root.DFSPreOrder() // Прямой обход
	fmt.Println()

	fmt.Println("DFS In-Order:")
	root.DFSInOrder() // Симметричный обход
	fmt.Println()

	fmt.Println("DFS Post-Order:")
	root.DFSPostOrder() // Обратный обход
	fmt.Println()
}