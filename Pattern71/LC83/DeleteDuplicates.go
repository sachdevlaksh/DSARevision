/*
LeetCode Problem #83: Remove Duplicates from Sorted List
Difficulty: Easy

Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.
*/

package LC83

///**
//* Definition for singly-linked list.
type ListNode struct {
     Val int
     Next *ListNode
 }

func deleteDuplicates(head *ListNode) *ListNode {
	copyHead := head
	for  copyHead != nil && copyHead.Next != nil{
		if copyHead.Val == copyHead.Next.Val{
			copyHead.Next = copyHead.Next.Next
		}else{
			copyHead = copyHead.Next
		}
	}
	return head
}