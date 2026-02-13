/*
LeetCode Problem #496: Next Greater Element I
Difficulty: Easy

The next greater element of some element x in an array is the first greater element to its right. You are given two distinct 0-indexed integer arrays nums1 and nums2, where nums1 is a subset of nums2.
*/

package LC496

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	m := make(map[int]int)
	stack := make([]int, 0)

	for i := len(nums2) - 1; i >= 0; i-- {
		x := nums2[i]
		for len(stack) > 0 && stack[len(stack)-1] <= x {
			stack = stack[:len(stack)-1]
		}
			if len(stack) == 0{
				m[x] = -1
			} else {
				m[x] = stack[len(stack)-1]
			}
			stack = append(stack, x)
	}
	outLi := make([]int, len(nums1))
	for i, v := range nums1{
		outLi[i] =  m[v]
	}

	return outLi
}
