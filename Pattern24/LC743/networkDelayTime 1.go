package LC743


import (
	"container/heap"
	"fmt"
	"math"
)

// Define a pair for storing {distance, node}
type pii struct {
	dist int
	node int
}

// PriorityQueue implements heap.Interface and holds pii.
type PriorityQueue []pii

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(pii))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func dijkstra(V int, adj [][]pii, S int) []int {
	dist := make([]int, V)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[S] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, pii{dist: 0, node: S})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(pii)
		d, u := curr.dist, curr.node

		if d > dist[u] {
			continue
		}

		for _, edge := range adj[u] {
			v, weight := edge.node, edge.dist
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(pq, pii{dist: dist[v], node: v})
			}
		}
	}

	return dist
}

func main() {
	V := 6
	adj := make([][]pii, V)

	adj[0] = append(adj[0], pii{node: 1, dist: 4})
	adj[0] = append(adj[0], pii{node: 2, dist: 1})
	adj[1] = append(adj[1], pii{node: 3, dist: 1})
	adj[2] = append(adj[2], pii{node: 1, dist: 2})
	adj[2] = append(adj[2], pii{node: 3, dist: 5})
	adj[3] = append(adj[3], pii{node: 4, dist: 3})
	adj[4] = append(adj[4], pii{node: 5, dist: 2})
	adj[5] = append(adj[5], pii{node: 3, dist: 1})

	sourceNode := 0

	shortestDistances := dijkstra(V, adj, sourceNode)

	fmt.Printf("Shortest distances from source node %d:\n", sourceNode)
	for i := 0; i < V; i++ {
		fmt.Printf("Node %d: %d\n", i, shortestDistances[i])
	}
}