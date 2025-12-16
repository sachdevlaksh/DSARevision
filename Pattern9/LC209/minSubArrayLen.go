package LC209

import "math"

func minSubArrayLen(target int, nums []int) int {
    mini := math.MaxInt
    sum := 0
	left := 0
    for right := 0; right < len(nums); right++ {
        sum += nums[right]
        for sum >= target {
            mini = min(mini, right-left+1)
            sum -= nums[left]
            left++
        }
    }
    if mini == math.MaxInt {
        return 0
    }
    return mini
}
