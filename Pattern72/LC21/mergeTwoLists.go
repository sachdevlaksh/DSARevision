/*
LeetCode Problem #21: Merge Two Sorted Lists
Difficulty: Easy

You are given the heads of two sorted linked lists list1 and list2. Merge the two lists into one sorted list and return the head of the merged list.
*/

package LC21
type ListNode struct {
     Val int
     Next *ListNode
 }
 
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil{
		return list2
	}

	if list2 == nil{
		return list1
	}

	 headNew := &ListNode{}

	out := headNew

	for list1 != nil && list2 != nil{
		if list1.Val <= list2.Val{
			out.Next = &ListNode{Val:list1.Val }
			list1 = list1.Next
		}else {
			out.Next = &ListNode{Val:list2.Val }
			list2 =  list2.Next
		}
		out=out.Next
	}
	if list1 != nil{
		out.Next = list1
	}else if list2 != nil{
		out.Next = list2
	}

	return headNew.Next
}