package LC143

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {

	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	curr := slow.Next
	slow.Next = nil
	var prev *ListNode

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	first,second := head, prev

	for second != nil{
		fNext := first.Next
		sNext :=  second.Next

		first.Next = second
		second.Next= fNext

		first= fNext
		second=sNext
	}

	return
}
