package LC202

func isHappy(n int) bool {

	m := make(map[int]bool)
	sum := 0
	for n != 1 {
		for n > 1 {
			rem := n % 10
			n = n / 10
			sum += rem * rem
			if sum == 1 {
				return true
			}
			if m[sum] {
				return false
			}
			m[sum] = true
			n = sum
		}
	}

	return true
}
