/*
LeetCode Problem #994: Rotting Oranges
Difficulty: Medium

You are given an m x n grid where each cell can have one of three values: 0 representing an empty cell, 1 representing a fresh orange, or 2 representing a rotten orange.
*/

package LC994

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	queue := make([][2]int, 0)

	head := 0
	fresh := 0

	for r, roV := range grid {
		for c, rocoV := range roV {
			if rocoV == 2 {
				queue = append(queue, [2]int{r, c})
			} else if rocoV == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	minutes := -1
	for head < len(queue) {
		levelSize := len(queue) - head
		minutes++
		for i := 0; i < levelSize; i++ {
			cell := queue[head]
			head++
            r, c := cell[0], cell[1]
			for _, d := range dirs {
				rd, cd := r+d[0], c+d[1]
				if rd >= 0 && cd >= 0 && rd < rows && cd < cols && grid[rd][cd] == 1 {
					grid[rd][cd] = 2
					fresh--
					queue = append(queue, [2]int{rd, cd})
				}
			}
		}

	}
	if fresh > 0 {
		return -1
	}
	return minutes
}

func orangesRottingV2(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	queue := make([][2]int, 0)
	head := 0
	fresh := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			} else if grid[r][c] == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	minutes := -1
	for head < len(queue) {
		levelSize := len(queue) - head
		minutes++
		for i := 0; i < levelSize; i++ {
			cell := queue[head]
			head++
			r, c := cell[0], cell[1]
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {

					grid[nr][nc] = 2
					fresh--
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
	}

	if fresh > 0 {
		return -1
	}
	return minutes

}
