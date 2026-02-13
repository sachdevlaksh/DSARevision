/*
LeetCode Problem #410: Split Array Largest Sum
Difficulty: Hard

Given an array nums which consists of non-negative integers and an integer m, you can split the array into m non-empty continuous subarrays. Write an algorithm to minimize the largest sum among these m subarrays.
*/

package LC410

func splitArray(nums []int, k int) int {
	l := 0
	r := 0
	n := len(nums)
	mx := 0
	for _, num := range nums {
		r += num
	}
	for l <= r {
		mid := l + (r-l)/2
		if isCorrectComb(nums, n, k, mid) {
			mx = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
		mid = l + (r-l)/2

	}

	return mx

}

func isCorrectComb(nums []int, n, m, mid int) bool {

	count := 1
	sum := 0
	for i:=0;i < n; i++{
		if nums[i] + sum <= mid{
			sum += nums[i]
		}else{
			count++
			if count > m || nums[i] > mid{
				return false
			}
			sum = nums[i]

		}
	}
	return true
}
