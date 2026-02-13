/*
LeetCode Problem #239: Sliding Window Maximum
Difficulty: Hard

You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window.
*/

package LC239


type DQ []int


func (m DQ) Len() int {
	return len(m)
}
func (dq *DQ) PopBack() {
	*dq = (*dq)[:len(*dq)-1]
}

func (dq *DQ) PopFront() {
	*dq = (*dq)[1:]
}
func (m *DQ) PushBack(x int) {
    *m = append(*m, x)
}

func (dq DQ) Front() int {
	return dq[0]
}

func (dq DQ) Back() int {
	return dq[len(dq)-1]
}
func maxSlidingWindow(nums []int, k int) []int {

	dq := &DQ{}
	n := len(nums)
	ans := make([]int, 0)

	for i := 0; i< n; i++{
		for dq.Len() != 0 && dq.Front() <= i-k{
			dq.PopFront()
		}

		for dq.Len() != 0 && nums[dq.Back()]  <= nums[i]{
			dq.PopBack()
		}

		dq.PushBack(i)

		if i >= k-1{
			ans= append(ans, nums[dq.Front()])
		}
	}

	return ans
}
