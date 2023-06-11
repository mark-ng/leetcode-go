// Day 4

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    moving := dummy
    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            moving.Next = list1
            list1 = list1.Next
        } else {
            moving.Next = list2
            list2 = list2.Next
        }
        moving = moving.Next
    }
    if list1 != nil {
        moving.Next = list1
    }
    if list2 != nil {
        moving.Next = list2
    }
    return dummy.Next
}