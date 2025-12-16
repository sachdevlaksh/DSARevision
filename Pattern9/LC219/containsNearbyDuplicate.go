package LC219

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int)

	for i, num := range nums{
		if idx, ok := seen[num]; ok && int(math.Abs(float64(idx-i))) <= k{
			return true
		}
		seen[num] = i
	}

	return false
}
func containsNearbyDuplicate2(nums []int, k int) bool {

	for left := 0; left < len(nums); left++{
		for right := left+1; right <= left + k && right < len(nums); right++{
			if nums[left] == nums[right]{
				return true
			}
		}
	}
	return false
}