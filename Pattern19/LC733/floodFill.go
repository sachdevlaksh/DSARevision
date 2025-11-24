package LC733

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	oriCol := image[sr][sc]
	if oriCol == color {
		return image
	}
	var dfs func(r, c int)

	dfs = func(r, c int) {
		if r >= 0 && r < len(image) && c >= 0 && c < len(image[0]) {
			if image[r][c] != oriCol {
				return
			}
			image[r][c] = color
			for _, d := range dirs {
				nr := r + d[0]
				nc := c + d[1]
				dfs(nr, nc)
			}
		}

	}
	dfs(sr, sc)
	return image

}
