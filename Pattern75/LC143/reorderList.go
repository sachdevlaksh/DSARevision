/*
LeetCode Problem #143: Reorder List
Difficulty: Medium

You are given the head of a singly linked-list. The list can be represented as: L0 → L1 → ... → Ln - 1 → Ln Reorder the list to be on the following form: L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → ...
*/

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
