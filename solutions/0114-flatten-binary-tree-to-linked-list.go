// Day 13

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func flatten(root *TreeNode)  {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        return
    }
    flatten(root.Right)

    flatten(root.Left)

    temp := root.Left
    root.Left = root.Right
    root.Right = temp

    p := root
    for p.Right != nil {
        p = p.Right
    }
    p.Right = root.Left
    root.Left = nil
}