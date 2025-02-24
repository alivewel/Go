package main

import (
	"fmt"
	"sync"
)

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

func (n *Node) Traversal(result *[]int, uniqueValues map[int]struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if n == nil {
		return
	}
	*result = append(*result, n.Data)
	uniqueValues[n.Data] = struct{}{}
	leftWg := sync.WaitGroup{}
	leftWg.Add(1)
	rightWg := sync.WaitGroup{}
	rightWg.Add(1)
	go n.Left.Traversal(result, uniqueValues, &leftWg)
	go n.Right.Traversal(result, uniqueValues, &rightWg)
	leftWg.Wait()
	rightWg.Wait()
}
	

// Написать функцию для обхода дерева в глубину
func (n *Node) DFS() {
	wg := sync.WaitGroup{}
	var result []int
	var uniqueValues = make(map[int]struct{}) 
	wg.Add(1)
	go n.Traversal(&result, uniqueValues, &wg)
	wg.Wait()
	sum := 0
	for value := range uniqueValues {
		sum += value
	}
	fmt.Println(result)
	fmt.Println(sum)
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
