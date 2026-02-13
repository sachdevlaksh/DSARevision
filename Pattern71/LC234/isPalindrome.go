/*
LeetCode Problem #234: Palindrome Linked List
Difficulty: Easy

Given the head of a singly linked list, return true if it is a palindrome or false otherwise.
*/

package LC234

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	l2 := reverseList(slow)
	l1 := head

	for l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return true
}

func reverseList(slow *ListNode) *ListNode {

	var next, prev *ListNode
	curr := slow
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
