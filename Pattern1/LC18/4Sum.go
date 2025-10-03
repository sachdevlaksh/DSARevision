package LC18

import "sort"
func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    out := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}
		num1 := nums[i]
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1  && nums[j] == nums[j-1]{
				continue
			}
			num2 := nums[j]
			l := j + 1
			r := len(nums) - 1
			for l < r {
				sum := num1 + num2 + nums[l] +nums[r]
				if sum == target {
					out = append(out, []int{num1, num2, nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum > target {
					r--
				} else {
					l++
				}
			}

		}

	}
	return out
}