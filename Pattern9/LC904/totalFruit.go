package LC904

func totalFruit(fruits []int) int {

	left := 0
	maxLen := 0
	ma := make(map[int]int)
	for right := 0;right < len(fruits);right++{
			ma[fruits[right]]++
			for len(ma) > 2{
				ma[fruits[left]]--
				if ma[fruits[left]] == 0{
					delete(ma, fruits[left])
				}
				left++
			}
			maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}