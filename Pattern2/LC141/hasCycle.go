/*
LeetCode Problem #141: Linked List Cycle
Difficulty: Easy

Given head, the head of a linked list, determine if the linked list has a cycle in it. There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
*/

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
