/*
LeetCode Problem #338: Counting Bits
Difficulty: Easy

Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.
*/

package LC338

func countBits(n int) []int {
    bits := make([]int, n+1)

    for i:= 0; i < n+1; i++{
        bits[i] = bits[i>>1] + i & 1
    }

    return bits[:n+1]
}