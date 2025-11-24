package main

import "sort"

func maxOperations(nums []int, k int) int {
	count := 0
	sort.Ints(nums)
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]+nums[j] < k {
			i++
		} else if nums[i]+nums[j] > k {
			j--
		} else {
			count++
			i++
			j--
		}
	}
	return count
}
