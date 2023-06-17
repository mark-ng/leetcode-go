// Day 8

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    ans := &ListNode{}
    ansMove := ans
    carry := 0
    for l1 != nil && l2 != nil {
        sum := l1.Val + l2.Val + carry
        if sum > 9 {
            sum = sum - 10
            carry = 1
        } else {
            carry = 0
        }
        ansMove.Next = &ListNode{
            Val: sum,
        }
        ansMove = ansMove.Next
        l1 = l1.Next
        l2 = l2.Next
    }
    for l1 != nil {
        sum := l1.Val + carry
        if sum > 9 {
            sum = sum - 10
            carry = 1
        } else {
            carry = 0
        }
        ansMove.Next = &ListNode{
            Val: sum,
        }
        ansMove = ansMove.Next
        l1 = l1.Next
    }
    for l2 != nil {
        sum := l2.Val + carry
        if sum > 9 {
            sum = sum - 10
            carry = 1
        } else {
            carry = 0
        }
        ansMove.Next = &ListNode{
            Val: sum,
        }
        ansMove = ansMove.Next
        l2 = l2.Next
    }
    if carry == 1 {
        ansMove.Next = &ListNode{
            Val: 1,
        }
    }
    return ans.Next
}