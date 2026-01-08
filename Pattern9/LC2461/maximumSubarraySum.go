package LC2461

func maximumSubarraySum(nums []int, k int) int64 {
	maxSum := int64(0)
	hm := make(map[int]int)
	sum := int64(0)
	left := 0
	for right := 0; right < len(nums); right++ {
		sum += int64(nums[right])
		hm[nums[right]]++
		for hm[nums[right]] > 1 || right-left+1 > k {
			hm[nums[left]]--
			sum -= int64(nums[left])
			left++
		}
		if right-left+1 == k {
			maxSum = max(maxSum, sum)
		}
	}

	return maxSum
}
