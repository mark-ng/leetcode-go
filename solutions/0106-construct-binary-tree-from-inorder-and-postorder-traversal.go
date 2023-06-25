// Day 14

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 || len(postorder) == 0 {
        return nil
    }
    if len(inorder) == 1 || len(postorder) == 1 {
        return &TreeNode{
            Val: postorder[len(postorder) - 1],
        }
    }
    // root
    root := postorder[len(postorder) - 1]

    // find index of root in inorder
    var idx int;
    for i := 0; i < len(inorder); i++ {
        if inorder[i] == root {
            idx = i
            break
        }
    }
    leftInorder := inorder[:idx] 
    rightInorder := inorder[idx + 1:]
    rightStartInclusive := len(postorder) - 1 - len(rightInorder)
    rightEndExclusive := rightStartInclusive + len(rightInorder)
    rightPostorder := postorder[rightStartInclusive : rightEndExclusive]
    leftPostorder := postorder[:rightStartInclusive]

    rootNode := TreeNode{
        Val: root,
        Left: buildTree(leftInorder, leftPostorder),
        Right: buildTree(rightInorder, rightPostorder),
    }
    return &rootNode
}