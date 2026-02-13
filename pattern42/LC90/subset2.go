/*
LeetCode Problem #90: Subsets II
Difficulty: Medium

Given an integer array nums that may contain duplicates, return all possible subsets (the power set). The solution set must not contain duplicate subsets.
*/

package LC90

import "sort"

func subsetsWithDup(nums []int) [][]int {

	var dfs func(int)
	var temp []int
	var result [][]int
	sort.Ints(nums)

	dfs = func(start int) {
		tempCopy := make([]int, len(temp))
		copy(tempCopy,temp)
		result = append(result, tempCopy)

		for j := start; j < len(nums); j++ {
			 if j> start && nums[j] == nums[j-1] {
				 continue
				}
				temp = append(temp, nums[j])
				dfs( j+1)
				temp = temp[:len(temp)-1]
		}

		return
	}

	dfs(0)

	return result
}