package LC2054

import (
	"container/heap"
	"sort"
)

type Event struct {
	end int
	val int
}

type MinHeap []Event

func (m MinHeap) Len() int {
	//TODO implement me
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	//TODO implement me
	return m[i].end < m[j].end
}

func (m MinHeap) Swap(i, j int) {
	//TODO implement me
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
	//TODO implement me
	*m = append(*m, x.(Event))
}

func (m *MinHeap) Pop() any {
	//TODO implement me
	out := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return out
}

func maxTwoEvents(events [][]int) int {

	sort.Slice(events, func(i, j int) bool {
	return 	events[i][0] < events[j][0]
	})

	minH := &MinHeap{}

	heap.Init(minH)
	maxPrev := 0
	ans := 0

	for _, eve := range events {
		startCurr, endCurr, valCurr := eve[0], eve[1], eve[2]

		for minH.Len() > 0 && startCurr > (*minH)[0].end {
			top := heap.Pop(minH).(Event)
			maxPrev = max(maxPrev, top.val)
		}
		ans = max(maxPrev+valCurr, ans)
		heap.Push(minH, Event{end: endCurr, val: valCurr})
	}

	return ans
}
