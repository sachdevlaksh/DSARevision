package LC713

func numSubarrayProductLessThanK(nums []int, k int) int {
	count := 0
	left := 0
	prod := 1
	for right := 0; right < len(nums); right++ {
		prod *= nums[right]
		for right >= left && prod >= k {
			prod = prod / nums[left]
			left++
		}
        count += right-left+1
	}
	return count
}
