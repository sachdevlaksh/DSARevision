package LC1658

func MinOperations(nums []int, x int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < x {
		return -1
	}

	target := sum - x

	if target == 0 {
		return len(nums)
	}
	curr := 0
	left := 0
	maxLen := -1
	for right := 0; right < len(nums); right++ {
		curr += nums[right]

		for curr > target {
			curr -= nums[left]
			left++
		}

		if curr == target {
			maxLen= max(maxLen, right-left+1)
		}
	}

	if maxLen == -1{
		return -1
	}

	return len(nums) - maxLen
}
