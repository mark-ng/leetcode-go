// Day 2

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    switch left, right := maxDepth(root.Left), maxDepth(root.Right); {
    case left >= right:
        return left + 1
    default:
        return right + 1
    }
}