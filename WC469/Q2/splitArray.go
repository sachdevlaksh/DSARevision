package Q2

import (
	"math"
)

func SplitArray(nums []int) int64 {

	if len(nums)  == 2{
		return int64(math.Abs(float64(nums[0] - nums[1])))
	}
	n := len(nums)
	breakOut := false
	j := -1
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] < nums[i-1]{
			return -1
		}
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			breakOut = true
			j = i
			break
		}
	}

	if !breakOut {
		return -1
	}
	for k :=j; k < n-1; {
		if nums[k+1] > nums[k] {
			return -1
		}
		k++
	}



	sl1 := calcSum(nums[:j])
	sl2 := calcSum(nums[j:])
	sl3 := calcSum(nums[:j+1])
	sl4 := calcSum(nums[j+1:])
	diff1 := math.Abs(float64(sl1 - sl2))
	diff2 := math.Abs(float64(sl3 - sl4))

	if diff1 < diff2 {
		return int64(diff1)
	}else {
		return 	int64(diff2)
	}
}

func calcSum(sl []int) (sum int64) {

	for i := 0; i < len(sl); i++ {
		sum += int64(sl[i])
	}
	return sum
}
