/*
LeetCode Problem #55: Jump Game
Difficulty: Medium

You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length from that position. Determine if you can reach the last index.
*/

package LC55

func canJump(nums []int) bool {

	maxReach := 0


	for i := 0; i < len(nums);i++{

		if i > maxReach{
			return false
		}

		if i + nums[i] > maxReach{
			maxReach = i+nums[i]
		}

		if maxReach >= len(nums) -1{
			return true
		}
	}

	return true

}