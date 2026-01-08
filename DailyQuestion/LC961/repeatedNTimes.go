package LC961

func repeatedNTimes(nums []int) int {

	n := len(nums) / 2
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
		if count[num] == n {
			return num
		}
	}

	return -1
}
