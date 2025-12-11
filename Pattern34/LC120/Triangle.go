package LC120

import "fmt"

func minimumTotalCopy(triangle [][]int) int {
	count := 0
	return getSum(triangle, 0, 0,count)
}
func getSum(triangle [][]int, i, j,count int) int {
     count =  count+1
	 fmt.Print( " Started for Func Index => ",count)
	if i == len(triangle) {
		fmt.Print( " Returned from Func Index => ",count)
		fmt.Print("  I => ", i)
		fmt.Println("  | J => ", j)
		return 0
	}

	current := triangle[i][j]
	fmt.Print(" | current => [", current)
	fmt.Print("]  |")
	fmt.Print( " Returned from Func Index => ",count)
	fmt.Print("  I => ", i)
	fmt.Println("  | J => ", j)
	left := getSum(triangle, i+1, j,count)
	right := getSum(triangle, i+1, j+1,count)
	sum := current + min(left, right)

	return sum
}

func MinimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle))

	for i, _ := range triangle {
		dp[i] = triangle[len(triangle)-1][i]
	}

	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			dp[col] = triangle[row][col] + min(dp[col], dp[col+1])
		}
	}

	return dp[0]
}

