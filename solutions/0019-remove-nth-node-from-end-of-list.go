// Day 12

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{}
    dummy.Next = head
    slow := dummy
    fast := dummy
    for i := 0; i < n + 1; i++ {
        fast = fast.Next
    }
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }
    // fmt.Println(slow.Val)
    slow.Next = slow.Next.Next
    return dummy.Next
}