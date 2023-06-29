/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
    ans := [][]int{}
    
    if root == nil {
        return ans 
    }

    type NodeDetail struct {
        node *TreeNode
        level int
    } 

    q := []NodeDetail{{root, 0}}

    curLevel := -1
    for len(q) != 0 {
        popNodeDetail := q[0]
        q = q[1:]

        ll := popNodeDetail.level
        popTreeNode := popNodeDetail.node

        if ll != curLevel {
            ans = append(ans, []int{})
            ans[ll] = append(ans[ll], popTreeNode.Val) 
            curLevel = ll
        } else {
            ans[ll] = append(ans[ll], popTreeNode.Val) 
        }

        if popTreeNode.Left != nil {
            q = append(q, NodeDetail{popTreeNode.Left, ll + 1})
        }
        if popTreeNode.Right != nil {
            q = append(q, NodeDetail{popTreeNode.Right, ll + 1})
        }
    }

    return ans
}