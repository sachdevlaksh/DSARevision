package LC3346

import "math"

func maxFrequency(nums []int, k int, numOperations int) int {

	count := 0

	for numOperations > 0{
		for _, num := range nums{
			if int(math.Abs(float64(float64(num)-float64(k)))) < k{
				count++
			}
		}
	}

	return count
}