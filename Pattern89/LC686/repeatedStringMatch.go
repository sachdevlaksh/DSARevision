/*
LeetCode Problem #686: Repeated String Match
Difficulty: Easy

Given two strings a and b, return the minimum number of times you should repeat string a so that string b is a substring of it. If b can never be a substring of a no matter how many times you repeat it, return -1.
*/

package LC686

func repeatedStringMatch(a string, b string) int {

	l1 := len(a) //4
	l2 := len(b) // 8
	or := a

	count := 1
	for l1 < l2 {
		count++
		a += or
		l1 = len(a)
	}

	for i := 0; i <= l1-l2; i++ {
		if b == a[i:i+l2] {
			return count
		}
	}

	count++
	a += or
	l1 = len(a)
	for i := 0; i <= l1-l2; i++ {
		if b == a[i:i+l2] {
			return count
		}
	}

	return -1
}
