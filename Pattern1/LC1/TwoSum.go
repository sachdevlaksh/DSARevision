/*
LeetCode Problem #1: Two Sum
Difficulty: Easy

Given an array of integers nums and an integer target, return the indices of the two numbers such that they add up to target. You may assume that each input would have exactly one solution, and you may not use the same element twice.
*/

package LC1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i, num := range nums {
		sec := target - num
		if index, ok := m[sec]; ok {
			return []int{i, index}
		}
		m[num] = i

	}
	return nil
}