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
