/*
LeetCode Problem #15: 3Sum
Difficulty: Medium

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
*/

package LC15

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	out := make([][]int, 0)
	for i, num := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicate nums[i]
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if num+nums[l]+nums[r] == 0 {
				out = append(out, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
			}

			if num+nums[l]+nums[r] < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return out
}
