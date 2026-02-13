/*
LeetCode Problem #42: Trapping Rain Water
Difficulty: Hard

Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
*/

package LC42

func trap(height []int) int {

	l := 0
	r := len(height) - 1
	lMax := 0
	rMax := 0
	total := 0

	for l < r {
		if height[l] < height[r] {
			if height[l] > lMax {
				lMax = height[l]
			} else {
				total += lMax - height[l]
			}
			l++
		} else {
			if height[r] > rMax {
				rMax = height[r]
			} else {
				total += rMax - height[r]
			}
			r--
		}
	}

	return total
}
