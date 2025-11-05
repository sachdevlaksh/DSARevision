package LC210

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)

	for _, gr := range prerequisites {
		graph[gr[1]] = append(graph[gr[1]], gr[0])
	}

	visited := make([]int, numCourses)
	stack := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			if !dfsOrder(visited, graph, i, &stack) {
				return []int{}
			}
		}
	}

	for l, m := 0, len(stack)-1; l < m; l, m = l+1, m-1 {
		stack[l], stack[m] = stack[m], stack[l]
	}

	return stack

}

func dfsOrder(visited []int, graph [][]int, node int, stack *[]int) bool {

	if visited[node] == 1 {
		return false
	}
	if visited[node] == 2 {
		return true

	}
	visited[node] = 1

	for _, adj := range graph[node] {
		if visited[adj] == 1 {
			// cycle detected
			return false
		}
		if visited[adj] == 0 {
			if !dfsOrder(visited, graph, adj, stack) {
				return false
			}
		}
	}

	visited[node] = 2
	*stack = append(*stack, node)
	return true
}
