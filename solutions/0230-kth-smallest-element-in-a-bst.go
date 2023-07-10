// Day 21

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func kthSmallest(root *TreeNode, k int) int {
    var ans int
    var count int = 0
    dfs(root, &count, k, &ans)
    return ans
}

func dfs(root *TreeNode, count *int, k int, ans *int) {
    if root == nil {
        return
    }
    
    dfs(root.Left, count, k, ans)
    *count = *count + 1
    if *count == k {
        *ans = root.Val
        return
    }
    dfs(root.Right, count, k, ans)
}