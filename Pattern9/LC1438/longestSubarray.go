package LC1438


func longestSubarray(nums []int, limit int) int {
	minLimit := -1
	for i := 0; i < len(nums);i++{
		if nums[i] <= limit{
			minLimit = 1
		}
		for j := i +1; j < len(nums); j++{
			if nums[j] - nums[i] < limit{
				return j-i
			}
		}
	}

	return minLimit
}