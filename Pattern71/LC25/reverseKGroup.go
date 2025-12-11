package LC25

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	if head == nil {
		return nil
	}


	var prev, next *ListNode
	curr := head
	count  := 0

	for curr != nil && count < k {
        curr = curr.Next
        count++
    }

    if count < k {
        // Not enough nodes to reverse, return head as is
        return head
    }
	for curr != nil && count < k{
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		count++
	}

	if next != nil{
		head.Next = reverseKGroup(next,k)
	}

	return prev
}


	func reverseKGroupPrac(head *ListNode, k int) *ListNode {

		if head == nil{
			return nil
		}
		count := 0
		curr := head
		for curr != nil && count < k{
			count++
			curr = curr.Next

		}
		if count < k{
			return head
		}

		curr = head
		count = 0
		var prev, next *ListNode

		for curr != nil && count < k{
			next = curr.Next
			curr.Next = prev
			prev = curr
			curr = next
			count++
		}
		if next != nil{
			head.Next = reverseKGroupPrac(next,k)
		}

		return prev

	}