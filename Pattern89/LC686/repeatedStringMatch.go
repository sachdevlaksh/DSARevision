package LC686

func repeatedStringMatch(a string, b string) int {

	l1 := len(a) //4
	l2 := len(b) // 8
	or := a

	count := 1
	for l1 < l2 {
		count++
		a += or
		l1 = len(a)
	}

	for i := 0; i <= l1-l2; i++ {
		if b == a[i:i+l2] {
			return count
		}
	}

	count++
	a += or
	l1 = len(a)
	for i := 0; i <= l1-l2; i++ {
		if b == a[i:i+l2] {
			return count
		}
	}

	return -1
}
