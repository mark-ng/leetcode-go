// Day 11

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sumNumbers(root *TreeNode) int {
    ans := 0
    dfs(root, 0, &ans)
    return ans
}

func dfs(root *TreeNode, cum int, ans *int) {
    if root.Left == nil && root.Right == nil {
        cum = cum * 10 + root.Val
        *ans += cum
        return 
    }
    if root.Left != nil {
        dfs(root.Left, cum * 10 + root.Val, ans)
    }
    if root.Right != nil {
        dfs(root.Right, cum * 10 + root.Val, ans)
    }
}