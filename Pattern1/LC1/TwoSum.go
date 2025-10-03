package LC1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i, num := range nums {
		sec := target - num
		if index, ok := m[sec]; ok {
			return []int{i, index}
		}
		m[num] = i

	}
	return nil
}