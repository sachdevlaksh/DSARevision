package LC40

import "slices"

func combinationSum2(candidates []int, target int) [][]int {
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
			// skip duplicates at the same recursion level
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			if sum+candidates[i] > target {
				// because sorted, no need to try larger elements
				break
			}
			// choose
			li = append(li, candidates[i])
			dfs(i+1, sum+candidates[i], li) // i+1 because each element can be used at most once
			// backtrack
			li = li[:len(li)-1]
		}
	}

	dfs(0, 0, make([]int, 0))
	return out
}
