package LC207

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
    indegree := make([]int, numCourses)

    // Build graph and indegree count
    for _, pre := range prerequisites {
        course, prereq := pre[0], pre[1]
        graph[prereq] = append(graph[prereq], course)
        indegree[course]++
    }

    // Queue for courses with no prerequisites
    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if indegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    count := 0
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
        count++

        for _, next := range graph[curr] {
            indegree[next]--
            if indegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }

    return count == numCourses
}

func CanFinish2(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
    for _, pre := range prerequisites {
        graph[pre[1]] = append(graph[pre[1]], pre[0])
    }

    visited := make([]int, numCourses)
    // visited: 0 = unvisited, 1 = visiting, 2 = visited

    var dfs func(int) bool
    dfs = func(course int) bool {
        if visited[course] == 1 { // Found a cycle
            return false
        }
        if visited[course] == 2 { // Already processed
            return true
        }

        visited[course] = 1 // Mark as visiting
        for _, next := range graph[course] {
            if !dfs(next) {
                return false
            }
        }
        visited[course] = 2 // Mark as done
        return true
    }

    for i := 0; i < numCourses; i++ {
        if !dfs(i) {
            return false
        }
    }
    return true
}

func CanFinish(numCourses int, prerequisites [][]int) bool {
    vis := make([]int, numCourses)
    adjac := make([][]int, numCourses)

    for _, val := range prerequisites{
        adjac[val[1]] = append(adjac[val[1]], val[0])
    }


    for i := 0; i< numCourses; i++{
        if vis[i] == 0{
            if !dfsItertative(i,vis,adjac){
                return false
            }
        }
    }

    return true
}

func dfsItertative(node int, vis []int, adj [][]int) bool  {
    if vis[node] == 1{
        return false
    }
    if vis[node] == 2{
        return true
    }
    vis[node] = 1

    for _, adjVer := range adj[node]{
        if !dfsItertative(adjVer,vis,adj){
            return false
        }
    }
    vis[node] = 2
    return true

}