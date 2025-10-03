package LC77

func Combine(n int, k int) [][]int {

	var result [][]int
	var temp []int

	var dfs func(int)

	dfs = func(start int) {
		if len(temp) == k{
			newSlice := make([]int, k)
			copy(newSlice,temp)
			result = append(result,newSlice)
			return
		}
		for i := start; i<= n; i++{
			temp = append(temp, i)
			dfs(i+1)
			temp = temp[:len(temp)-1]
		}
		return
	}
	dfs(1)

	 return result
}
