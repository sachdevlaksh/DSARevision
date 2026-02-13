/*
LeetCode Problem #344: Reverse String
Difficulty: Easy

Write a function that reverses a string. The input string is given as an array of characters s. You must do this by modifying the input array in-place with O(1) extra memory.
*/

package LC344

func reverseString(s []byte)  {

	n := len(s)
	if n <= 1{
		return
	}
	left := 0
	right := n-1

	for left<= right{
		s[right],s[left] = s[left],s[right]
		left++
		right--
	}
}