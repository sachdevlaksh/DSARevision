/*
LeetCode Problem #20: Valid Parentheses
Difficulty: Easy

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid. An input string is valid if: Open brackets must be closed by the same type of brackets, and open brackets must be closed in the correct order.
*/

package LC20

func IsValid(s string) bool {

	if len(s) % 2 == 1{
        return false
    }
	m1 := make(map[rune]rune)
	m1['{'] = '}'
	m1['['] = ']'
	m1['('] = ')'
	m2 := make(map[rune]rune)
	m2['}'] = '{'
	m2[']'] = '['
	m2[')'] = '('

	li := make([]rune, 0)
	li = append(li,rune(s[0]))
	for i := 1; i < len(s); i++ {
		if _, ok := m1[rune(s[i])]; ok {
			li = append(li, rune(s[i]))
		} else {
			if m2[rune(s[i])] == (li[len(li)-1]) {
				li = li[:len(li)-1]
			}else{
				return false
			}
		}
	}

	if len(li) != 0{
		return false
	}
	return true

}
