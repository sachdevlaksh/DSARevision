/*
LeetCode Problem #23: Merge k Sorted Lists
Difficulty: Hard

You are given an array of k linked-lists lists, each linked-list is sorted in ascending order. Merge all the linked-lists into one sorted linked-list and return it.
*/

package LC23

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type minHeap []*ListNode

func (m minHeap) Len() int {
	//TODO implement me
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	//TODO implement me
	return m[i].Val < m[j].Val
}

func (m minHeap) Swap(i, j int) {
	//TODO implement me
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x any) {
	//TODO implement me
	*m = append(*m, x.(*ListNode))
}

func (m *minHeap) Pop() any {
	//TODO implement me
	last := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return last
}

func mergeKLists(lists []*ListNode) *ListNode {

	mi := &minHeap{}
	heap.Init(mi)

	for _, list := range lists {
		if list != nil {
			heap.Push(mi,list)
		}
	}

	dummy := &ListNode{}
	out := dummy

	for mi.Len() > 0 {
		node := heap.Pop(mi).(*ListNode)
		out.Next = node
		out = out.Next
		if node.Next != nil {
			heap.Push(mi,node.Next)
		}
	}

	return dummy.Next
}
