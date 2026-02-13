/*
LeetCode Problem #48: Rotate Image
Difficulty: Medium

You are given an n x n 2D matrix representing an image. Rotate the image by 90 degrees (clockwise).
*/

package LC48

func RotateImage(matrix [][]int)  {
    n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

    for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}
