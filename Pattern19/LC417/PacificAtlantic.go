/*
LeetCode Problem #417: Pacific Atlantic Water Flow
Difficulty: Medium

There is an m x n rectangular island that borders both the Pacific and Atlantic oceans. The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.
*/

package LC417

func PacificAtlantic(heights [][]int) [][]int {

	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}
	rows := len(heights)
	cols := len(heights[0])
	paci := make([][]bool, rows)
	atlan := make([][]bool, rows)

	for i := range paci {
		paci[i] = make([]bool, cols)
		atlan[i] = make([]bool, cols)
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(int, int, [][]bool)

	dfs = func(r int, c int, mat [][]bool) {

		if mat[r][c] {
			return
		}
		mat[r][c] = true

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nc >= 0 && nr < rows && nc < cols && heights[nr][nc] >= heights[r][c] {
				dfs(nr, nc, mat)
			}
		}
	}

	for r := 0; r < rows; r++ {
		dfs(r, 0, paci)
		dfs(r, cols-1, atlan)
	}

	for c := 0; c < cols; c++ {
		dfs(0, c, paci)
		dfs(rows-1, c, atlan)
	}
	out := [][]int{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if paci[i][j] && atlan[i][j] {
				out = append(out, []int{i, j})
			}
		}
	}
	return out
}
func pacificAtlantic(heights [][]int) [][]int {

	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	rows := len(heights)
	cols := len(heights[0])
	paci := make([][]bool, rows)
	atlan := make([][]bool, rows)
	for i := range paci {
		paci[i] = make([]bool, cols)
		atlan[i] = make([]bool, cols)
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var dfs func(int, int, [][]bool)
	dfs = func(r, c int, ocean [][]bool) {
		if ocean[r][c] {
			return
		}
		ocean[r][c] = true

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nc >= 0 && nr < rows && nc < cols && heights[nr][nc] >= heights[r][c] {
				dfs(nr, nc, ocean)
			}
		}
	}

	// Run DFS from Pacific border
	for r := 0; r < rows; r++ {
		dfs(r, 0, paci)       // Left column
		dfs(r, cols-1, atlan) // Right column
	}
	for c := 0; c < cols; c++ {
		dfs(0, c, paci)       // Top row
		dfs(rows-1, c, atlan) // Bottom row
	}

	out := [][]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if paci[r][c] && atlan[r][c] {
				out = append(out, []int{r, c})
			}
		}
	}

	return out
}
