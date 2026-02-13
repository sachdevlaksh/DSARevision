/*
LeetCode Problem #167: Two Sum II - Input Array Is Sorted
Difficulty: Medium

Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number.
*/

package LC167

func twoSum(numbers []int, target int) []int {
    l, r := 0, len(numbers) - 1

    for l < r {
        if numbers[l] + numbers[r] == target{
            return []int{l+1,r+1}
        }else if numbers[l] + numbers[r] < target{
            l++
        }else{
            r--
        }
    }

    return []int{}
}