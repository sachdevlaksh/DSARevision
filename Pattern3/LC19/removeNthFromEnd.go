package LC19

type ListNode struct {
    Val  int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

    if head == nil {
        return nil
    }
    countNode := head
    leng := 0
    for countNode != nil {
        leng++
        countNode = countNode.Next
    }

    if n == leng {
        return head.Next
    }

    count := leng - n - 1
    dummy := head

    for count > 0 {
        dummy = dummy.Next
        count--

    }
    if dummy.Next != nil {
        dummy.Next = dummy.Next.Next
    }

    return head
}
