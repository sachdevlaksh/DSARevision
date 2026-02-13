/*
LeetCode Problem #547: Number of Provinces
Difficulty: Medium

There are n cities. Some of them are connected, while some are not. If city a is connected directly with city b, and city b is connected directly with city c, then city a is connected indirectly with city c.
*/

package LC547

func findCircleNum(isConnected [][]int) int {
	if len(isConnected) == 0 {
		return 0
	}
	n := len(isConnected)
	vis := make([]bool, n)
	count := 0
	for i := 0; i < n; i++ {
		if !vis[i] {
			count++
			bfs(isConnected, &vis, i)
		}
	}
	return count

}

func bfs(isConnected [][]int, v *[]bool, i int) {
	(*v)[i] = true
	que := make([]int, 0)
	que = append(que, i)

	for len(que) > 0 {
		node := que[0]
		que = que[1:]
		for k := 0; k < len(isConnected); k++ {

			if isConnected[node][k] == 1 && !(*v)[k] {
				(*v)[k] = true
				que = append(que, k)
			}
		}
	}
}
