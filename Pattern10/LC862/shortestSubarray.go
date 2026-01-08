package LC862

import "math"

func ShortestSubarray(nums []int, k int) int {
	n := len(nums)
	mini := math.MaxInt
	prefix := make([]int, n+1)
	dq := make([]int, 0)

	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	for i := 0; i <= n; i++ {
		for len(dq) > 0 && prefix[i]-prefix[dq[0]] >= k {
			mini = min(mini, i-dq[0])
			dq = dq[1:]
		}
		for len(dq) > 0 && prefix[i] <= prefix[dq[len(dq)-1]] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
	}

	if mini == math.MaxInt {
		return -1
	}

	return mini
}
