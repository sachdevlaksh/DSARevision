/*
LeetCode Problem #506: Relative Ranks
Difficulty: Easy

You are given an integer array score of size n where score[i] is the score of the ith athlete in a competition. All the scores are guaranteed to be unique. The athletes are placed based on their scores, where the athlete with the highest score gets first place, the athlete with the second highest score gets second place, and so on.
*/

package LC506

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Score int
	Index int
}

func findRelativeRanks(score []int) []string {
	pq := make(MinHeap, len(score))
	result := make([]string, len(score))

	for i, sc := range score {
		pq[i] = &Item{Score: sc, Index: i}
	}
	heap.Init(&pq)
	for i := 0; pq.Len() > 0; i++ {
		score := heap.Pop(&pq).(*Item)
		switch i {
		case 0:
			result[score.Index] = "Gold Medal"
		case 1:
			result[score.Index] = "Silver Medal"
		case 2:
			result[score.Index] = "Bronze Medal"
		default:
			result[score.Index] = fmt.Sprintf("%d", i+1)

		}
	}

	return result
}

type MinHeap []*Item

func (m MinHeap) Len() int {
	//TODO implement me
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	//TODO implement me
	return m[i].Score > m[j].Score

}

func (m MinHeap) Swap(i, j int) {
	//TODO implement me
	m[i], m[j] = m[j], m[i]

}

func (m *MinHeap) Push(x any) {
	//TODO implement me
	*m = append(*m, x.(*Item))

}

func (m *MinHeap) Pop() any {
	//TODO implement me
	length := m.Len()
	old := *m
	item := old[length-1]
	*m = old[:length-1]
	return item

}
