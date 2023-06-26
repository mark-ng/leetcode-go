// Day 15

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    res := &ListNode{}
    ans := res
    repeat := 101
    for head != nil {
        if head.Next != nil {
            if head.Val != head.Next.Val && head.Val != repeat {
                ans.Next = head
                ans = ans.Next
            } else if head.Val == head.Next.Val {
                repeat = head.Val
            }
        } else {
            if head.Val != repeat {
                ans.Next = head
                ans = ans.Next
            } else {
                ans.Next = nil
            }
        }
        head = head.Next
    }
    return res.Next
}