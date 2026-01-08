package LC695

func maxAreaOfIsland(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	n := len(grid)
	m := len(grid[0])

	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}

	maxArea := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !vis[i][j] && grid[i][j] == 1 {
				area := 1
				area = bfs(grid, &vis, i, j, n, m, area)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea

}

func bfs(grid [][]int, v *[][]bool, i int, j int, n int, m, area int) int {

	(*v)[i][j] = true
	que := make([][2]int, 0)
	que = append(que, [2]int{i, j})
	for len(que) > 0 {
		roww := que[0][0]
		coll := que[0][1]
		que = que[1:]

		dRow := []int{-1, 0, 1, 0}
		dCol := []int{0, 1, 0, -1}
		for k := 0; k < 4; k++ {
			newR := roww + dRow[k]
			newC := coll + dCol[k]
			if newR >= 0 && newR < n && newC >= 0 && newC < m && grid[newR][newC] == 1 && !(*v)[newR][newC] {
				(*v)[newR][newC] = true
				que = append(que, [2]int{newR, newC})
				area++
			}
		}
	}
	return area
}
