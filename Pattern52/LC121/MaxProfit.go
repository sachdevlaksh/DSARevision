package LC121

import "math"


func maxProfit(prices []int) int {
	n := len(prices)

	maxProf := 0
	minPrice := math.MaxInt

	for i := 0; i < n; i++{
		minPrice = min(minPrice, prices[i])
		maxProf = max(maxProf, prices[i]-minPrice)
	}

	return maxProf
}