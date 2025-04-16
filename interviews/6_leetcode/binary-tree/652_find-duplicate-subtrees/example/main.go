type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    var result []*TreeNode
    count := make(map[string]int)

    var dfs func(node *TreeNode) string
    dfs = func(node *TreeNode) string {
        if node == nil {
            return "#"
        }
        serial := fmt.Sprintf("%d,%s,%s", node.Val, dfs(node.Left), dfs(node.Right))
        count[serial]++
        if count[serial] == 2 {
            result = append(result, node)
        }
        return serial
    }

    dfs(root)
    return result
}
