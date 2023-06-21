// Day 10

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return mirror(root.Left, root.Right)
}

func mirror(t1 *TreeNode, t2 *TreeNode) bool {
    if t1 == nil && t2 == nil {
        return true
    }
    if t1 == nil || t2 == nil {
        return false
    }

    return t1.Val == t2.Val && mirror(t1.Left, t2.Right) && mirror(t1.Right, t2.Left)
}