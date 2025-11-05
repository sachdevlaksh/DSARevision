package LC84

func largestRectangleArea(heights []int) int {
	n := len(heights)
	maxArea := 0
	next := getNextSmallElement(heights, n)
	prev := getPrevSmallElement(heights, n)
	for i := 0; i < n; i++ {
		if next[i] == -1 {
			next[i] = n
		}
		area := heights[i] * (next[i] - prev[i] - 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func getNextSmallElement(heights []int, n int) []int {
	stacks := make([]int, 0)
	ans := make([]int, n)
	stacks = append(stacks, -1)
	for i := n - 1; i >= 0; i-- {
		for len(stacks) > 1 &&  heights[i] <= heights[top(&stacks)] {
			pop(&stacks)
		}
		ans[i] = top(&stacks)
		push(&stacks, i)
	}
	return ans
}
func getPrevSmallElement(heights []int, n int) []int {
	stacks := make([]int, 0)
	ans := make([]int, n)
	stacks = append(stacks, -1)
	for i := 0; i < n; i++ {
		for len(stacks) > 1 &&  heights[i] <= heights[top(&stacks)] {
			pop(&stacks)
		}
		ans[i] = top(&stacks)
		push(&stacks, i)
	}
	return ans
}
func defStacks(nums []int) *[]int {
	stacks := make([]int, len(nums))
	copy(stacks, nums)

	return &stacks
}

func pop(stacks *[]int) int {
	old := *stacks
	n := len(old)
	*stacks = (*stacks)[:n-1]
	return old[n-1]
}

func top(stacks *[]int) int {
	return (*stacks)[len(*stacks)-1]
}
func push(stacks *[]int, ele int) *[]int {

	*stacks = append(*stacks, ele)
	return stacks
}
