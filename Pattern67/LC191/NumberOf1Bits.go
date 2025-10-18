package LC191

func hammingWeight(n int) int {
	count := 0
	for n > 0 {
		count += int(n & 1)
		n >>= 1
	}

	return count
}