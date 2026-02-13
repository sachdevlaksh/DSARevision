/*
LeetCode Problem #191: Number of 1 Bits
Difficulty: Easy

Write a function that takes the binary representation of an unsigned integer and returns the number of '1' bits it has.
*/

package LC191

func hammingWeight(n int) int {
	count := 0
	for n > 0 {
		count += int(n & 1)
		n >>= 1
	}

	return count
}