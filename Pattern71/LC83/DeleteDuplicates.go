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