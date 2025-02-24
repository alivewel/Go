package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTree(nodes []interface{}, index int) *TreeNode {
	if index >= len(nodes) || nodes[index] == nil {
		return nil
	}

	node := &TreeNode{Val: nodes[index].(int)}
	node.Left = createTree(nodes, 2*index+1)
	node.Right = createTree(nodes, 2*index+2)
	return node
}

func main() {
	nodes1 := []interface{}{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}
	nodes2 := []interface{}{1, 2, 3, 4, 5, 6, 7}
	root1 := createTree(nodes1, 0)
	root2 := createTree(nodes2, 0)

	// Call the preorderTraversal function
	result := isSameTree(root1, root2)
	fmt.Println(result) // Output the result
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil 
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
