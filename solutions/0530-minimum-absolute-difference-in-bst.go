// Day 17

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// var minDiff int = -1
// var last int = -1

func getMinimumDifference(root *TreeNode) int {
    last := -1
    minDiff := 100001
    inorder(root, &last, &minDiff)
    return minDiff
}

func inorder(root *TreeNode, last *int, minDiff *int) {
    if root == nil {
        return
    }
    inorder(root.Right, last, minDiff)
    if *last != -1 {
        diff := *last - root.Val
        if diff < *minDiff {
            *minDiff = diff
        }
    }
    *last = root.Val
    inorder(root.Left, last, minDiff)
}