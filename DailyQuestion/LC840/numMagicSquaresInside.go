/*
LeetCode Problem #840: Magic Squares In Grid
Difficulty: Medium

A 3 x 3 magic square is a 3 x 3 grid filled with distinct numbers from 1 to 9 such that each row, column, and both diagonals all have the same sum.
*/

package LC840

func numMagicSquaresInside(grid [][]int) int {

	n := len(grid)
	m := len(grid[0])
	for n < 3 || m < 3 {
		return 0
	}
	out := 0

	for i := 0; i <= n-3; i++ {
		for j := 0; j <= m-3; j++ {
			if checkMagicSquare(i, j, grid) {
				out++
			}
		}
	}
	return out
}

func checkMagicSquare(i, j int, nums [][]int) bool {

	seen := make(map[int]bool)
	for r := i; r < i+3; r++ {
		for c := j; c < j+3; c++ {
			v := nums[r][c]
			if v < 1 || v > 9 || seen[v] {
				return false
			}
			seen[v] = true
		}
	}
	d := []int{0, 1, 2}
	prevSum := 0
	for c := j; c < j+3; c++ {
		prevSum += nums[i][c]
	}


	//rows
	for k := 0; k < 3; k++ {
		r := i + d[k]
		sum := 0
		for k1 := 0; k1 < 3; k1++ {
			c := j + d[k1]
			sum += nums[r][c]
		}
		if sum != prevSum {
			return false
		}

	}
	for k := 0; k < 3; k++ {

		c := j + d[k]
		sum := 0
		for k1 := 0; k1 < 3; k1++ {
			r := i + d[k1]
			sum += nums[r][c]
		}
		if sum != prevSum {
			return false
		}
	}
	sum := 0
	for k := 0; k < 3; k++ {
		c := j + d[k]
		r := i + d[k]
		sum += nums[r][c]
	}
	if sum != prevSum {
		return false
	}

	sum = 0
	d1 := []int{2, 1, 0}
	for k := 0; k < 3; k++ {
		c := j + d[k]
		r := i + d1[k]
		sum += nums[r][c]
	}
	if sum != prevSum {
		return false
	}
	return true
}
