// Day 9

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func averageOfLevels(root *TreeNode) []float64 {
    q := []*TreeNode{}
    q = append(q, root)
    nextChild := 1
    childCount := 0

    sum := 0
    count := 0
    ans := []float64{}

    for len(q) != 0 {
        for i := 0; i < nextChild; i++ {
            popNode := q[0]
            sum += popNode.Val
            count += 1
            q = q[1:]

            if popNode.Left != nil {
                q = append(q, popNode.Left)
                childCount++
            }
            if popNode.Right != nil {
                q = append(q, popNode.Right)
                childCount++
            }
        }

        ans = append(ans, float64(sum) / float64(count))
        sum = 0
        count = 0

        nextChild = childCount
        childCount = 0
    }
    return ans
}