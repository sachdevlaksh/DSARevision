/*
LeetCode Problem #60: Permutation Sequence
Difficulty: Hard

The set [1, 2, 3, ..., n] contains a total of n! unique permutations. By listing and labeling all of the permutations in order, we get the following sequence for n = 3. Given n and k, return the kth permutation sequence.
*/

package LC60

import (
	"strconv"
	"strings"
)

func getPermutation(n int, k int) string {
	fact := 1
	nums := make([]int, 0)
	var sb strings.Builder
	for i := 1; i < n; i++ {
		fact *= i
		nums = append(nums, i)
	}
	nums = append(nums, n)
	k = k - 1
	for {
		sb.WriteString(strconv.Itoa(nums[k/fact]))
		nums = append(nums[:k/fact], nums[k/fact+1:]...)
		if len(nums) == 0 {
			break
		}
		k = k % fact
		fact = fact / len(nums)
	}

	return sb.String()
}
