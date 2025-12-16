package LC424

func characterReplacement(s string, k int) int {
	maxF,left,res := 0,0,0
	li := make([]int, 26)
	for right := 0; right < len(s); right++{
		li[s[right]-'A']++
		maxF = max(maxF, li[s[right]-'A'])
		for right-left+1 > k + maxF{
			li[s[left]-'A']--
			left++
		}
		res = max(res,right-left+1)
	}
	return res
}