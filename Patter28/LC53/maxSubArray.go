package LC53

import "math"

func MaxSubArray(nums []int) int {
	n := len(nums)
	maxSum := math.MinInt
	sum := 0

	if n == 1 {
		return nums[0]
	} else if n == 0 {
		return 0
	}

	for i := 0; i < n; i++ {
		sum += nums[i]
		maxSum = max(sum, maxSum)
		if sum < 0 {
			sum = 0
		}

	}

	return maxSum
}

func MaxSubArrayPrint(nums []int) []int {
	n := len(nums)
	maxSum := math.MinInt
	sum := 0

	if n <= 1 {
		return nums
	}

	start, end := -1, -1
	for i := 0; i < n; i++ {
		if sum == 0 {
			start = i
		}
		sum += nums[i]
		if sum > maxSum{
			end = i
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}

	}
	return nums[start:end+1]
}