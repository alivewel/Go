// You can edit this code!
// Click here and start typing.
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

var (
   Sum map[int]struct{}
   M sync.Mutex
)

func NewNode(data int) *Node {

	node := new(Node)

	node.Data = data
	node.Left = nil
	node.Right = nil

	return node
}

// в глубину
func (n *Node) DFS(wg *sync.WaitGroup) {
	fmt.Println(n.Data)
	Sum[n.Data] = struct{}{}
	if n.Left != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n.Left.DFS(wg)
		}()
	}
	if n.Right != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n.Right.DFS(wg)
		}()
	}
}

func main() {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right.Left = NewNode(6)
	root.Right.Right = NewNode(7)

	wg := sync.WaitGroup{}
	Sum = make(map[int]struct{})
	
	//wg.Add(1)
	root.DFS(&wg)
	wg.Wait()
	
	var sums int  = 0
	for k := range Sum {
		sums += k
	}
	fmt.Println(sums)
}
