// Day 5

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

 func copyRandomList(head *Node) *Node {
    headhead := head
    headdump := head

    // Create an array to hold the new linkedlist
    _, cloneArr := clone(head)

    cur := 0
    for headdump != nil {
        rand := headdump.Random
        idx := findNode(headhead, rand)
        if idx == 10002 {
            headdump = headdump.Next
            cur++
            continue
        }
        if idx == 10001 {
            cloneArr[cur].Random = nil 
        } else {
            cloneArr[cur].Random = cloneArr[idx]

        }
        headdump = headdump.Next
        cur++
    }
    if len(cloneArr) > 0 {
        return cloneArr[0]
    } else {
        return nil
    }
}

// find random node index in original linkedlist
func findNode(head *Node, random *Node) int {
    if random == nil {
        return 10001
    }
    i := 0
    for head != nil {
        if (head == random) {
            return i
        }
        head = head.Next
        i++
    }
    return 10002
}

// Return the start of the linkedlist, return the array
func clone(head *Node) (*Node, []*Node) {
    clone := &Node{}
    movingClone := clone
    arr := []*Node{}
    for head != nil {
        // Create a new Node
        newNode := &Node{
            Val: head.Val,
        }

        //Add the arr
        arr = append(arr, newNode)

        // Add to list
        movingClone.Next = newNode
        movingClone = movingClone.Next

        head = head.Next
    }
    return clone.Next, arr
}
