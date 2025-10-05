package Q1

func AlternatingSum(nums []int) int {

	sum := 0
	for i := 0; i< len(nums); i= i+2{
		sum += nums[i]
	}
	for j := 1; j< len(nums); j= j+2{
		sum -= nums[j]
	}

	return sum

}