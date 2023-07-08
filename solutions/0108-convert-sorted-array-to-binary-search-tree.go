// Day 20

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sortedArrayToBST(nums []int) *TreeNode {
    return dfs(nums, 0, len(nums))
}
func dfs(nums []int, l int, r int) *TreeNode {
    if l >= r {
        return nil
    } 
    mid := (l + r) / 2
    node := TreeNode{}
    node.Val = nums[mid]
    node.Left = dfs(nums, l, mid)
    node.Right = dfs(nums, mid + 1, r)
    return &node 
}