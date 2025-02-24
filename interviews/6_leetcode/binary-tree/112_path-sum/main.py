class Solution:
    def isLeaf(self, node: TreeNode):
        return node.left is None and node.right is None
    
    # time: 0(n)
    # mem: 0(n)
    # note: можно не заводить доп. переменную, а делать
    # targetSum = targetSum - node.val и сравнивать targetSum с 0
    def hasSum(self, node: TreeNode, currSum: int, targetSum: int) -> bool:
        # currSum - префиксный массив
        if node is None:
            return False
        # если лист и дает нужную сумму, то мы нашли ответ
        if self.isLeaf(node) and node.val + currSum == targetSum: 
            return True
        isLeftSubTreeHasSum = self.hasSum(node.left, currSum + node.val, targetSum) 
        isRightSubTreeHasSum = self.hasSum(node.right, currSum + node.val, targetSum) 
        return isLeftSubTreeHasSum or isRightSubTreeHasSum

    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool: 
        return self. hasSum(root, 0, targetSum)
