package LC3

func LengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}
	left := 0
	maxV := 0

	seen := make(map[byte]int)

	seen[s[0]] = 0

	for right := 0; right < len(s); right++ {
		if idx, ok := seen[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		seen[s[right]] = right
		maxV= max(maxV, right-left+1)
	}

	return maxV
}
