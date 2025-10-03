package LC1493

func LongestSubarray(nums []int) int {
	l, r, out, isZero := 0, 0, 0, false

	for _, num := range nums {
		if num == 1 {
			r++
		} else {
			l = r
			r = 0
			isZero = true
		}
		out = max(out, l+r)

	}
	if isZero {
		return out
	} else {
		return out - 1
	}

}
