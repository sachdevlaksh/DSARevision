package LC621

func leastInterval(tasks []byte, n int) int {

	l := make([]int, 26)

	for _, t := range tasks {
		l[t-'A']++
	}

	maxFreq := 0
	maxFreqCount := 0

	for _, l1 := range l {
		if l1 > maxFreq {
			maxFreqCount = 1
			maxFreq = l1
		} else if l1 == maxFreq {
			maxFreqCount++
		}
	}

	res := (maxFreq-1)*(n+1) + maxFreqCount

	if res < len(tasks) {
		return len(tasks)
	} else {
		return res
	}

}
