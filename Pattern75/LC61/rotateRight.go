/*
LeetCode Problem #61: Rotate List
Difficulty: Medium

Given the head of a linked list, rotate the list to the right by k places.
*/

package LC61

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 || head.Next == nil {
		return head
	}

	length := 1
	tail := head

	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	k = k % length
	if k == 0 {
		return head
	}

	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	slow := head
	for fast.Next!=nil{
		slow= slow.Next
		fast=fast.Next
	}

	newHead := slow.Next
	slow.Next=nil
	fast.Next=head

	return newHead

}
