package LC743

import (
	"container/heap"
	"math"
)

func NetworkDelayTime(times [][]int, n int, k int) int {

	adjList := make([][]networkDelay, n+1)
	for _, time := range times {
		adjList[time[0]] = append(adjList[time[0]], networkDelay{dist: time[2], node: time[1]})
	}

	return dijks(adjList, n, k)
}

func dijks(adjList [][]networkDelay, n int, k int) int {
	dist := make([]int, n+1)

	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	minPq := &pq{}
	heap.Init(minPq)
	heap.Push(minPq, networkDelay{dist: 0, node: k})

	for minPq.Len() > 0 {
		curr := heap.Pop(minPq).(networkDelay)
		d, nod := curr.dist, curr.node

		if d > dist[nod] {
			continue
		}

		for _, adj := range adjList[nod] {
			adjNode, adjWeight := adj.node, adj.dist

			if dist[nod]+adjWeight < dist[adjNode] {
				dist[adjNode] = dist[nod] + adjWeight
				heap.Push(minPq, networkDelay{node: adjNode, dist: dist[adjNode]})
			}
		}

	}
	maxDelay := 0
	for i:=1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1
		}
		if dist[i] > maxDelay {
			maxDelay = dist[i]
		}
	}

	return maxDelay

}

type pq []networkDelay

func (p pq) Len() int {
	//TODO implement me
	return len(p)
}

func (p pq) Less(i, j int) bool {
	//TODO implement me
	return p[i].dist < p[j].dist
}

func (p pq) Swap(i, j int) {
	//TODO implement me
	p[i], p[j] = p[j], p[i]
}

func (p *pq) Push(x any) {
	//TODO implement me
	*p = append(*p, x.(networkDelay))
}

func (p *pq) Pop() any {
	//TODO implement me
	leng := len(*p)
	lengEl := (*p)[leng-1]
	*p = (*p)[:leng-1]
	return lengEl
}

type networkDelay struct {
	dist int
	node int
}
type MinHeap []networkDelay
