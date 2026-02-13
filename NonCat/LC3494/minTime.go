/*
LeetCode Problem #3494: Maximum Sum of Non-Overlapping Subarrays With Constraint
Difficulty: Hard

Given an array nums and an integer minK, find the maximum sum of a subarray while satisfying the following constraint: Every element in the subarray must be greater than or equal to minBit.
*/

package LC3494

import "math"

type Solution struct{}

func (s *Solution) minTime(skill []int, mana []int) int64 {
	n := len(skill)
	m := len(mana)
	finishTime := make([]int64, n)

	for i := 0; i < m; i++ {
		now := finishTime[0]
		currPotion := int64(mana[i])

		for j := 1; j < n; j++ {
			now = int64(math.Max(float64(now+int64(skill[j-1])*currPotion), float64(finishTime[j])))
		}
		finishTime[n-1] = now + int64(skill[n-1])*currPotion

		for j := n - 2; j >= 0; j-- {
			finishTime[j] = finishTime[j+1] - int64(skill[j+1])*currPotion
		}
	}

	return finishTime[n-1]
}

func MinTime(skill []int, mana []int) int64 {

	n := len(skill)
	m := len(mana)
	prev := make([]int64, n)
	var out int64
	prev[0]= int64(mana[0])
	for i := 1; i < n; i++ {
		prev[i] = prev[i-1]+int64(skill[i] * mana[0])
	}
	out = prev[n-1]
	for i := 1; i < m; i++ {

		sum := int64(0)
		l := int64(0)
		r := out

		for l < r {

			mid := (l + r) / 2

			for j := 0; j < n-2; j++ {
				temp := mid + int64(skill[j] * mana[i])
				if prev[i+1] < temp {
					l = mid + 1
					continue
				} else if prev[i+1] > temp {
					prev[i] = temp
					sum += temp
				}
			}
		}
		out += sum
	}
	return out
}
