/*
LeetCode Problem #19: Remove Nth Node From End of List
Difficulty: Medium

Given the head of a linked list, remove the nth node from the end of the list and return its head.
*/

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
