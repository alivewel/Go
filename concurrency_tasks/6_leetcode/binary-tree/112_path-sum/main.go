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
	// nodes1 := []interface{}{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}
	nodes := []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}
	root := createTree(nodes, 0)

	// Call the preorderTraversal function
	result := hasPathSum(root, 22)
	fmt.Println(result) // Output the result
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return checkSum(root, root.Val, targetSum)
}

func checkSum(root *TreeNode, curSum, targetSum int) bool {
	if root == nil { // очень важное условие - без него паника.
		return false
	}
	if isLeaf(root) && targetSum == curSum+root.Val {
		return true
	}
	resLeft := checkSum(root.Left, curSum+root.Val, targetSum)
	resRight := checkSum(root.Right, curSum+root.Val, targetSum)
	return resLeft || resRight
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}
