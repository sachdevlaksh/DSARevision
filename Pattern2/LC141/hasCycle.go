package LC141

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

func hasCycle2(head *ListNode) bool {

	hs := make(map[*ListNode]bool)

	for head != nil {
		if seen := hs[head]; seen {
			return true
		}
		hs[head] = true
	}

	return false
}
