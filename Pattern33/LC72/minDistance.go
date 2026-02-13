/*
LeetCode Problem #72: Edit Distance
Difficulty: Hard

Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2. You have three operations permitted on a word: Insert a character, Delete a character, Replace a character.
*/

package LC72

func minDistance(word1 string, word2 string) int {

	//return solve(a,b, 0,0)

	dp := make([][]int, len(word1))

	for i := range dp {
		dp[i] = make([]int, len(word2))
	}
	for i, v := range dp {
		for j := range v {
			dp[i][j] = -1
		}
	}
	return solveMem(word1, word2, 0, 0, &dp)

}

func solve(as string, bs string, i, j int) int {

	if i > len(bs) {
		return 1+ len(as) - i
	}
	if j > len(as) {
		return 1+ len(bs) - i
	}

	if as[i] == bs[i] {
		return 1+ solve(as, bs, i+1, j+1)
	}else{
		//min(insert,replace,delete
		return min(solve(as, bs, i, j+1), solve(as, bs, i+1, j+1), solve(as, bs, i+1, j))
	}
}
func solveMem(as string, bs string, i, j int, dp *[][]int) int {

	if i == len(as) {
		return len(bs) - j
	}
	if j == len(bs) {
		return len(as) - i
	}

	if (*dp)[i][j] != -1 {
		return (*dp)[i][j]
	}

	if as[i] == bs[j] {
		return solveMem(as, bs, i+1, j+1, dp)
	} else {
		//min(insert,replace,delete
		(*dp)[i][j] = min(solveMem(as, bs, i, j+1, dp)+1, solveMem(as, bs, i+1, j+1, dp)+1, solveMem(as, bs, i+1, j, dp)+1)
		return (*dp)[i][j]
	}
}
