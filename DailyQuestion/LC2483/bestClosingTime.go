/*
LeetCode Problem #2483: Minimum Penalty for a Shop
Difficulty: Medium

You are given the closing time for a shop, represented by a string s of length 10 in the format 'HH:MM'. You want to find the latest time when the shop should close such that the penalty is minimized.
*/

package LC2483

import "math"

func bestClosingTime(customers string) int {

	n := len(customers)
	p := make([]int, n+1)
	s := make([]int, n+1)
	out := -1
	mini := math.MaxInt
	for i := 0; i < n; i++ {
		p[i+1] = p[i]
		if customers[i] == 'N' {
			p[i+1]++
		}

		j := n - 1 - i
		s[j] = s[j+1]
		if customers[j] == 'Y' {
			s[j]++
		}

	}

	for i := 0; i <= n; i++ {
		if mini > p[i]+s[i] {
			mini = p[i] + s[i]
			out = i
		}
	}
	return out
}

func bestClosingTimeOptimized(customers string) int {
	var n int = len(customers)
	var pen = 0
	var mini = 0
	var result = n
	for indx := n - 1; indx >= 0; indx-- {
		if customers[indx] == 'Y' {
			pen++
		} else {
			pen--
		}
		if pen <= mini {
			mini = pen
			result = indx
		}
	}
	return result
}
func MaximumScore(nums []int) int64 {
	n := len(nums)

	// Step 1: build suffixMin array
	suffixMin := make([]int, n)
	suffixMin[n-1] = nums[n-1]

	for i := n - 2; i >= 0; i-- {
		if nums[i] < suffixMin[i+1] {
			suffixMin[i] = nums[i]
		} else {
			suffixMin[i] = suffixMin[i+1]
		}
	}

	// Step 2: iterate splits
	var prefixSum int64 = 0
	var maxScore int64 = math.MinInt64

	for i := 0; i < n-1; i++ {
		prefixSum += int64(nums[i])
		score := prefixSum - int64(suffixMin[i+1])
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}

func maximumScoreEasy(nums []int) int64 {
	n := len(nums)

	p := make([]int64, n+1) // prefix sums
	s := make([]int, n+1)   // suffix minimums

	// initialize suffix minimum sentinel
	s[n] = math.MaxInt

	// build prefix sum and suffix min together (same loop style)
	for i := 0; i < n; i++ {
		// prefix sum
		p[i+1] = p[i] + int64(nums[i])

		// suffix minimum
		j := n - 1 - i
		if i == 0 {
			s[j] = nums[j]
		} else {
			if nums[j] < s[j+1] {
				s[j] = nums[j]
			} else {
				s[j] = s[j+1]
			}
			s[j] = min(s[j+1],nums[j] )
		}
	}

	var maxScore int64 = math.MinInt64

	// valid splits: 0 <= i < n-1
	for i := 0; i < n-1; i++ {
		score := p[i+1] - int64(s[i+1])
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}