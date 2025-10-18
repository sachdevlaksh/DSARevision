package LC66
func PlusOne(nums []int) []int {
	n := len(nums)

	carry := 1
	for i := n - 1; i >= 0; i-- {
		if nums[i] == 9 && carry == 1 {
			nums[i] = 0
			carry = 1
			continue
		}
		nums[i] = nums[i] + carry
		carry = 0
	}
	if nums[0] == 0 {
		nums = append([]int{1}, nums...)
		return nums
	}
	return nums
}