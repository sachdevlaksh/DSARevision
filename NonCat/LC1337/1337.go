/*
LeetCode Problem #1337
Title: Problem 1337
Difficulty: Unknown

LeetCode Problem #1337
*/

package LC1337

import (
	"cmp"
	"container/heap"
	"slices"
	"sort"
)

type Item struct {
	strength int
	posIn    int
}

func kWeakestRows(mat [][]int, k int) []int {
	h := &MaxHeap{}
	heap.Init(h)
	for i, row := range mat {
		stength := 0
		for _, val := range row {
			stength += val
		}
		heap.Push(h, Item{strength: stength, posIn: i})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	slices.SortFunc(*h, func(a, b Item) int {
		return cmp.Or(cmp.Compare(a.strength, b.strength),
			cmp.Compare(a.posIn,b.posIn))
	})

	sort.Slice(*h, func(i, j Item) bool {
    if i.posIn != j.posIn {
        return i.posIn > j.posIn
    }
    return i.strength > j.strength
	
	result := make([]int, h.Len())
	for i := 0; i < h.Len(); i++ {
		result[i] = (*h)[i].posIn
	}
	return result
}

type MaxHeap []Item

func (m MaxHeap) Len() int {
	//TODO implement me
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	//TODO implement me
	return m[i].strength < m[j].strength
}

func (m MaxHeap) Swap(i, j int) {
	//TODO implement me
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	//TODO implement me
	*m = append(*m, x.(Item))

}

func (m *MaxHeap) Pop() any {
	//TODO implement me
	old := *m
	n := len(old)
	item := old[n-1]
	*m = old[:n-1]
	return item

}
