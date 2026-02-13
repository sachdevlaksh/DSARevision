/*
LeetCode Problem #13: Roman to Integer
Difficulty: Easy

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M. Given a roman numeral, convert it to an integer.
*/

package LC13

func romanToInt(s string) int {
	m := make(map[rune]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return m[rune(s[0])]
	}

	out := 0
	for i := 0; i < len(s)-1; i++ {
		if m[rune(s[i])] < m[rune(s[i+1])] {
			out -= m[rune(s[i])]
		} else {
			out += m[rune(s[i])]
		}
	}

	return out + m[rune(s[len(s)-1])]
}
