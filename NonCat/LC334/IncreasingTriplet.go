package LC334

import (
	"fmt"
	"math"
)

func IncreasingTriplet(nums []int) bool {

	small, mid := math.MaxInt32, math.MaxInt32

	for _, num := range nums {
		if num <= small  {
			small = num
		} else if num < mid  {
			mid = num
		} else {
			fmt.Println(true)
			return true

		}
	}
	fmt.Println(false)
	return false
}
