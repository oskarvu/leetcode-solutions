/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    node := &ListNode{0, nil}
    head := node
    pass := 0
    currentL1 := l1
    currentL2 := l2
    for currentL1 != nil && currentL2 != nil {
        sum := currentL1.Val + currentL2.Val + pass
        pass = sum / 10
        node.Val = sum % 10
        if currentL1.Next != nil || currentL2.Next != nil {
            node.Next = &ListNode{0, nil}
        } else if pass != 0 {
            node.Next = &ListNode{pass, nil}
        } else {
            node.Next = nil
        }
        node = node.Next
        currentL1 = currentL1.Next
        currentL2 = currentL2.Next
    }
    var current *ListNode
    if currentL1 == nil {
        current = currentL2
    } else {
        current = currentL1
    }
    for current != nil {
        sum := current.Val + pass
        pass = sum / 10
        node.Val = sum % 10
        if current.Next != nil {
            node.Next = &ListNode{0, nil}
        } else if pass != 0 {
            node.Next = &ListNode{pass, nil}
        } else {
            node.Next = nil
        }
        node = node.Next
        current = current.Next
    }
    return head
}