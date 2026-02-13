/*
LeetCode Problem #2300: Successful Pairs of Spells and Potions
Difficulty: Medium

You are given two positive integer arrays spells and potions, of length n and m respectively, where spells[i] represents the strength of the ith spell and potions[j] represents the strength of the jth potion.
*/

package LC2300

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {

	n := len(spells)
	m := len(potions)
	sort.Ints(potions)
	spells2 := make([]int, n)
	sort.Ints(spells2)
	copy(spells2, spells)
	ma := make(map[int64]int, n)

	for _, s := range spells2 {
		for j, p := range potions {
			if int64(p)*int64(s) < success {
				continue
			}
			ma[int64(s)] = m - j
			break
		}
	}
	for i, v := range spells {
		spells[i] = ma[int64(v)]
	}
	return spells
}

func SuccessfulPairs(spells []int, potions []int, success int64) []int {

	m := len(potions)
	sort.Ints(potions)
	for i, s := range spells {
		lookMultiplier := (success + int64(s) - 1)/int64(s)
            l,r := 0, m
			for l < r {
                mid := (l+r)/2
				if int64(potions[mid]) < lookMultiplier{
					l = mid + 1
				}else{
					r = mid
            }
		}
		spells[i] = m -l

	}
	return spells
}