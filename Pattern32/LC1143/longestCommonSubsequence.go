/*
LeetCode Problem #1143: Longest Common Subsequence
Difficulty: Medium

Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.
*/

package LC1143

func longestCommonSubsequence(text1 string, text2 string) int {

	return compare1(text1,text2, 0,0)

}

func compare1(t1, t2 string, i, j int) int{

	if i >= len(t1) || j >= len(t2){
		return 0
	}
	if t1[i] == t2[j]{
		return 1 + compare1(t1,t2,i+1,j+1)
	}else{
		return max(compare1(t1,t2,i+1,j),compare1(t1,t2,i,j+1))
	}
}


func LongestCommonSubsequence(t1 string, t2 string) int {

	dp := make([][]int, len(t1))


	for i := range dp {
		dp[i] = make([]int, len(t2))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return compare2(t1,t2, 0,0, &dp)

}

func compare2(t1, t2 string, i, j int, dp *[][]int) int{

	if i >= len(t1) || j >= len(t2){
		return 0
	}

	if (*dp)[i][j] != -1{
		return (*dp)[i][j]
	}
	if t1[i] == t2[j]{
		(*dp)[i][j] = 1 + compare2(t1,t2,i+1,j+1, dp)
		return (*dp)[i][j]
	}else{
		return max(compare2(t1,t2,i+1,j, dp),compare2(t1,t2,i,j+1, dp))
	}
}
