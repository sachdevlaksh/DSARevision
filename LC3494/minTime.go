package LC3494

import "math"

package main

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
	var out int64
	for i := 0; i < n; i++ {
		out += int64(skill[0] * mana[i])
	}

	for i := 1; i < n; i++ {

		l := int64(0)
		r := out
		sum := int64(0)
		for l < r {
			prev := make([]int64, n)
			mid := (l + r) / 2

			for j := 0; i < m-1; j++ {
				temp := int64(skill[i] * mana[j])
				if out < temp {
					l = mid + 1
				} else {
					prev[i] = temp
					sum += temp
					r = mid
				}
			}
		}
		out += sum
	}
	return out
}
