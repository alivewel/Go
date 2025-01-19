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
	// nodes := []interface{}{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}
	nodes := []interface{}{1, 2, 3, 4, 5, 6, 7}
	root := createTree(nodes, 0)

	// Call the preorderTraversal function
	result := inorderTraversal(root)
	fmt.Println(result) // Output the result
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	traverse(root, &result)
	return result
}

func traverse(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	traverse(root.Left, result)
	*result = append(*result, root.Val)
	traverse(root.Right, result)
}
