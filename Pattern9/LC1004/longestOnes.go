package LC1004

func longestOnes(nums []int, k int) int {
	maxLeng := 0
	zeroLen := 0
	left := 0
	for right := 0; right < len(nums);right++{
		if nums[right] == 0{
			zeroLen++
		}
		for zeroLen > 0{
			if nums[right] == 0{
				zeroLen--
			}
			left++

		}
		maxLeng = max(maxLeng, right-left+1)
	}
	return maxLeng
}