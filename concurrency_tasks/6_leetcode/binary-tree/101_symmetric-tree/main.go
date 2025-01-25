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
	result := levelOrder(root)
	fmt.Println(result) // Output the result
}

func isSymmetric(root *TreeNode) bool {
	
	return true 
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	return preOrder(root, 0, &result)
}

func preOrder(root *TreeNode, level int, result *[][]int) [][]int {
	if root == nil {
		return *result
	}
	if level == len(*result) {
		*result = append(*result, []int{})
	}
	(*result)[level] = append((*result)[level], root.Val)
	*result = preOrder(root.Left, level+1, result)
	*result = preOrder(root.Right, level+1, result)
	return *result
}
