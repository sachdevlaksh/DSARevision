/*
LeetCode Problem #39: Combination Sum
Difficulty: Medium

Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target.
*/

package LC39

import "slices"
func combinationSum(candidates []int, target int) [][]int {
	out := make([][]int, 0)
	slices.Sort(candidates)
	var dfs func(index, sum int, li []int)

	dfs = func(index, sum int, li []int) {
		if sum == target {
			clone := make([]int, len(li))
			copy(clone, li)
			out = append(out, clone)
			return
		}
		if sum > target {
			return
		}
		for i := index; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				// because sorted, no need to try larger elements
				break
			}
			// choose
			li = append(li, candidates[i])
			dfs(i, sum+candidates[i], li) // i+1 because each element can be used at most once
			// backtrack
			li = li[:len(li)-1]
		}
	}

	dfs(0, 0, make([]int, 0))
	return out
}