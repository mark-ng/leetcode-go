// Day 19

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    State []int
    Pos int
}

func Constructor(root *TreeNode) BSTIterator {
    state := []int{}
    pos := 0 // always point to next
    dfs(root, &state) 
    return BSTIterator{
        State: state,
        Pos: pos,
    }
}

func dfs(root *TreeNode, arr *[]int) {
    if root == nil {
        return
    }
    dfs(root.Left, arr)
    *arr = append(*arr, root.Val)
    dfs(root.Right, arr)
}


func (this *BSTIterator) Next() int {
    result := this.State[this.Pos]
    this.Pos++
    return result
}


func (this *BSTIterator) HasNext() bool {
    return this.Pos < len(this.State)
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */