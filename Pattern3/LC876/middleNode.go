/*
LeetCode Problem #876: Middle of the Linked List
Difficulty: Easy

Given the head of a singly linked list, return the middle node of the linked list. If there are two middle nodes, return the second middle node.
*/

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
