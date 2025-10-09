package LC1046

import "container/heap"

func lastStoneWeight(stones []int) int {

	maxHeap := MaxHeap{}
	heap.Init(&maxHeap)
	for _, val := range stones {
		heap.Push(&maxHeap, val)
	}

	for maxHeap.Len() > 1 {
		f := heap.Pop(&maxHeap).(int)
		s := heap.Pop(&maxHeap).(int)
		if s != f {
			heap.Push(&maxHeap, f-s)
		}
	}
	return heap.Pop(&maxHeap).(int)
}

type MaxHeap []int

func (m MaxHeap) Len() int {

	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {

	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {

	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {

	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() any {

	length := m.Len()
	old := *m
	item := old[length-1]
	*m = old[:length-1]
	return item
}
