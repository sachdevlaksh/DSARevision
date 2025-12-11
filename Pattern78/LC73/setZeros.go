package LC73



func setZeroes(matrix [][]int) {
	ro := len(matrix)
	co := len(matrix[0])
	list := make([][2]int, 0)
	for i := 0; i < ro; i++ {
		for j := 0; j < co; j++ {
			if matrix[i][j] == 0 {
	 		list = append(list, [2]int{i,j})
			}
		}
	}
	for _, v := range list{
		solve(&matrix, v[0], v[1], ro, co)
	}
}

func solve(matrix *[][]int, ri, ci, ro, co int) {

	for i := 0; i < ro; i++ {
		(*matrix)[i][ci] = 0
	}
	for i := 0; i < co; i++ {
		(*matrix)[ri][i] = 0
	}
}
