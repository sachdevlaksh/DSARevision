/*
LeetCode Problem #1970: Last Day Where You Can Still Cross
Difficulty: Hard

There is a 1-based binary matrix where 0 represents water and 1 represents land. Find the last day you can cross from the top to the bottom by walking only on land cells.
*/

package LC1970
func latestDayToCross(row int, col int, cells [][]int) int {

	var dfs func(i, j int, vis [][]int, water [][]int) bool

	dfs = func(i, j int, vis, water [][]int) bool {
		if i < 1 || j < 1 || i > row || j > col || vis[i][j] == 1 || water[i][j] == 1 {
			return false
		}
		if i == row {
			return true
		}

		vis[i][j] = 1

		return dfs(i+1, j, vis, water) ||
			dfs(i-1, j, vis, water) ||
			dfs(i, j+1, vis, water) ||
			dfs(i, j-1, vis, water)
	}

	canCross := func(day int) bool {
		water := make([][]int, row+1)
		for i := range water {
			water[i] = make([]int, col+1)
		}
		for i := 0; i <= day; i++ {
			water[cells[i][0]][cells[i][1]] = 1
		}

		for c := 1; c <= col; c++ {
			vis := make([][]int, row+1)
			for i := 0; i <= row; i++ {
				vis[i] = make([]int, col+1)
			}
			if dfs(1, c, vis, water) {
				return true
			}
		}
		return false
	}

	left, right := 0, len(cells)-1
	ans := 0
	for left <= right {
		mid := left + (right-left)/2
		if canCross(mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func latestDayToCross2(row int, col int, cells [][]int) int {

	var dfs func(i, j int, vis [][]int, dp [][]int, water [][]int) bool

	dfs = func(i, j int, currVis, dp [][]int, water [][]int) bool {
		if i < 1 || j < 1 || i > row || j > col || currVis[i][j] == 1 || water[i][j] == 1 {
			return false
		}
		if i == row {
			return true
		}

		if dp[i][j] != -1 {
			return dp[i][j] == 1
		}

		currVis[i][j] = 1

		canReach :=
			dfs(i+1, j, currVis, dp, water) ||
				dfs(i-1, j, currVis, dp, water) ||
				dfs(i, j+1, currVis, dp, water) ||
				dfs(i, j-1, currVis, dp, water)

		if canReach {
			dp[i][j] = 1
		} else {
			dp[i][j] = 0
		}
		return canReach
	}

	canCross := func(day int) bool {
		// rebuild water grid up to `day`
		water := make([][]int, row+1)
		for i := range water {
			water[i] = make([]int, col+1)
		}
		for i := 0; i <= day; i++ {
			r, c := cells[i][0], cells[i][1]
			water[r][c] = 1
		}

		dp := make([][]int, row+1)
		for i := 0; i <= row; i++ {
			dp[i] = make([]int, col+1)
			for j := 0; j <= col; j++ {
				dp[i][j] = -1
			}
		}

		for c := 1; c <= col; c++ {
			vis := make([][]int, col+1)
			if dfs(1, c, vis, dp, water) {
				return true
			}
		}
		return false
	}

	// ---- Binary Search (ONLY new logic) ----
	left, right := 0, len(cells)-1
	ans := 0

	for left <= right {
		mid := left + (right-left)/2
		if canCross(mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}
