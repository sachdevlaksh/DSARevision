/*
LeetCode Problem #215: Kth Largest Element in an Array
Difficulty: Medium

Given an integer array nums and an integer k, return the kth largest element in the array. Note that it is the kth largest element in the sorted order, not the kth distinct element.
*/

package LC215

import "container/heap"


func findKthLargest(nums []int, k int) int {
	h := MaxHeap(nums[:k])
	heap.Init(&h)
	for _,num := range nums[k:]{
		if num > h[0]{
			heap.Pop(&h)
			heap.Push(&h,num)
		}

	}
	return h[0]
}

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i],h[j] = h[j],h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h,x.(int))
}

func (h *MaxHeap) Pop() any {
	length := len(*h)
	popEle := (*h)[length-1]
	*h = (*h)[:length-1]
	return popEle
}
