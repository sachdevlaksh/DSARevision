/*
LeetCode Problem #202: Happy Number
Difficulty: Easy

Write an algorithm to determine if a number n is happy. A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits. Repeat until the number equals 1 or it loops endlessly in a cycle.
*/

package LC202

func isHappy(n int) bool {

	m := make(map[int]bool)
	sum := 0
	for n != 1 {
		for n > 1 {
			rem := n % 10
			n = n / 10
			sum += rem * rem
			if sum == 1 {
				return true
			}
			if m[sum] {
				return false
			}
			m[sum] = true
			n = sum
		}
	}

	return true
}
