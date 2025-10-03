package LC78

func Subsets(nums []int) [][]int {

	
	var dfs func(int)
	var temp []int
	var result [][]int

	dfs = func(start int) {
		tempCopy := make([]int, len(temp))
		copy(tempCopy,temp)
		result = append(result, tempCopy)

		for j := start; j < len(nums); j++ {
				temp = append(temp, nums[j])
				dfs( j+1)
				temp = temp[:len(temp)-1]
		}

		return
	}

	dfs(0)

	return result
}
