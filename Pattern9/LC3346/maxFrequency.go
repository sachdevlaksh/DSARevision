/*
LeetCode Problem #3346: Find the Maximum Length of Valid Roll Sequence
Difficulty: Hard

A sequence of rolls is said to be valid if for every pair of consecutive rolls, the absolute difference is between 2 and n inclusive.
*/

package LC3346

import "math"

func maxFrequency(nums []int, k int, numOperations int) int {

	count := 0

	for numOperations > 0{
		for _, num := range nums{
			if int(math.Abs(float64(float64(num)-float64(k)))) < k{
				count++
			}
		}
	}

	return count
}