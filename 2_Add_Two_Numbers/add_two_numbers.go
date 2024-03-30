package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var carry int
    dummy := ListNode{
        Val:  -1,
        Next: nil,
    }
    current := &dummy

    for l1 != nil || l2 != nil || carry != 0 {
        sum := 0
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        sum += carry

        carry = sum / 10
        sum %= 10

        current.Next = &ListNode{
            Val:  sum,
            Next: nil,
        }
        current = current.Next
    }

    return dummy.Next
}
// l1 = [2,4,3], l2 = [5,6,4]
//
