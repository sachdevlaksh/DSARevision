/*
LeetCode Problem #31: Next Permutation
Difficulty: Medium

Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
*/

package LC31

import "fmt"

func NextPermutation(nums []int) {

	if len(nums) <= 1{
		return
	}
	i := len(nums) - 2

	for i >= 0 && nums[i] >= nums[i+1]{
		i--
	}

	if i >= 0{

		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i]{
			j--
		}

		nums[i],nums[j]= nums[j],nums[i]
	}
	reverse(nums, i+1,len(nums)-1)
	fmt.Print(nums)
}

func reverse(nums[]int, i , j int){
	for i < j{
		nums[i],nums[j] = nums[j],nums[i]
		i++
		j--
	}
}