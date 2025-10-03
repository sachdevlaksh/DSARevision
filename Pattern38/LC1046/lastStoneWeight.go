package LC1046

import "container/heap"

func lastStoneWeight(stones []int) int {

	maxHeap := MaxHeap2{}
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

type MaxHeap2 []int

func (m MaxHeap2) Len() int {

	return len(m)
}

func (m MaxHeap2) Less(i, j int) bool {

	return m[i] > m[j]
}

func (m MaxHeap2) Swap(i, j int) {

	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap2) Push(x any) {

	*m = append(*m, x.(int))
}

func (m *MaxHeap2) Pop() any {

	length := m.Len()
	old := *m
	item := old[length-1]
	*m = old[:length-1]
	return item
}
