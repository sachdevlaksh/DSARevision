/*
LeetCode Problem #904: Fruit Into Baskets
Difficulty: Medium

You are visiting a farm that has unlimited rows of fruit trees arranged in a straight line. Each row has a name label by an integer. You have two baskets and each basket can carry unlimited amount of fruit, but you want each basket to only carry one type of fruit.
*/

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