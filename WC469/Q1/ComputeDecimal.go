package Q1

import (
	"math"
	"strconv"
)

func DecimalRepresentation(n int) []int {
	out := make([]int, 0)
	for n > 0 {
		str := strconv.Itoa(n)
		strLen := len(str)
		div1 := math.Pow(10, float64(strLen-1))
		out = append(out, n/int(div1)*int(math.Pow(10, float64(strLen-1))))
		n = n % int(div1)
	}

	return out

}
