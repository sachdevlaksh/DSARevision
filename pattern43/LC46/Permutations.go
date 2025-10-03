package LC46

func Permute(nums []int) [][]int {
	var ans [][]int
	 permutation(&nums, &ans, 0)
	return ans
}

func permutation(nums *[]int, ans *[][]int, i int) {

	if i >= len(*nums) {
		*ans = append(*ans, append([]int{}, *nums...))
		return
	}

	for j:= i; j < len(*nums); j++{
		(*nums)[i],(*nums)[j] = (*nums)[j],(*nums)[i]
		permutation(nums,ans,i+1)
		(*nums)[i],(*nums)[j] = (*nums)[j],(*nums)[i]
	}

	return

}
