package main

import "container/heap"

type KthLargest struct {
	MaxHeap *MinHeap
	Limit   int
}

func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)
	for _, ele := range nums {
		heap.Push(h, ele)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return KthLargest{h, k}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.MaxHeap, val)

	if this.MaxHeap.Len() > this.Limit {
		heap.Pop(this.MaxHeap)
	}
	return this.MaxHeap.Min().(int)
}

type MinHeap []int

func (m MinHeap) Len() int {
	//TODO implement me
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	//TODO implement me
	return m[i]<m[j]
}

func (m MinHeap) Swap(i, j int) {
	//TODO implement me
	m[i],m[j]=m[j],m[i]
}

func (m *MinHeap) Push(x any) {
	//TODO implement me
	*m =append(*m,x.(int))
}

func (m *MinHeap) Pop() (v any) {
	//TODO implement me
	v, *m = (*m)[len(*m)-1],(*m)[:len(*m)-1]
	return
}
func (h *MinHeap) Min() (v any) {
	return (*h)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
