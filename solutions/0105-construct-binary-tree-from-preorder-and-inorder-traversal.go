// Day 12

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }

    root := preorder[0] 
    var rootPos int
    for i := 0; i < len(inorder); i++ {
        if inorder[i] == root {
            rootPos = i
            break
        }
    }
    leftTree := inorder[:rootPos]
    rightTree := inorder[rootPos + 1:]

    leftTreeLen := len(leftTree)

    n := TreeNode{
        Val: root, 
    
    return &n
}