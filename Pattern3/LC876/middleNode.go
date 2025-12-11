package LC876

type ListNode struct {
    Val  int
    Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
  	countNode := head
	leng := 0
	for countNode != nil {
		leng++
		countNode = countNode.Next
	}
	steps := 0
	if leng%2 == 1 {
		steps = leng / 2
	} else {
		steps = leng / 2
	}


	for steps > 0 {
		head = head.Next
        steps--
	}

	return head
}
