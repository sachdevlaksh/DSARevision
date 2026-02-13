/*
LeetCode Problem #78: Subsets
Difficulty: Medium

Given an integer array nums of unique elements, return all the possible subsets (the power set). The solution set must not contain duplicate subsets. Return the solution in any order.
*/

package LC78

func Subsets(nums []int) [][]int {
res := [][]int{}
    for k:=0; k<=len(nums); k++{
      find(nums, 0, k, []int{}, &res)
    }
    return res
}

func find(nums []int, idx, k int, list []int, res *[][]int){
    if len(list) == k{
        temp := make([]int, k)
        copy(temp, list)
        *res = append(*res, temp)
        return
    }

    for i := idx; i< len(nums); i++{
        find(nums, i+1, k, append(list, nums[i]), res)
    }
 }
