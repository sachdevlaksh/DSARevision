package LC75

func sortColors(nums []int) {

	left, mid, right := 0, 0, len(nums)-1

	for mid <= right {
		switch nums[mid] {

		case 0:
			nums[left], nums[mid] = nums[mid], nums[left]
			left++
			mid++
		case 1:
			mid++
		case 2:
			nums[right], nums[mid] = nums[mid], nums[right]
			right--

		}
	}
}
func sortColors2(nums []int) {
	count := [3]int{}

	for _, v := range nums {
		count[v]++
	}

	idx := 0
	for color := 0; color <= 2; color++ {
		for count[color] > 0 {
			nums[idx] = color
			idx++
			count[color]--
		}
	}
}
