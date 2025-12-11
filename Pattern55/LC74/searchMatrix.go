package LC74

func searchMatrix(matrix [][]int, target int) bool {

	li := make([]int, 0)

	for i := range matrix {
		li = append(li, matrix[i][0])
	}

	le, ri := 0, len(li)-1
	mid := -1
	for le <= ri {
		mid = (le + ri) / 2
		if li[mid] > target {
			ri = mid - 1
			continue
		}
		if li[mid] < target {
			le = mid + 1
			continue
		}
		return true
	}

	if ri < 0 {
		return false
	}
	liNew := matrix[ri]
	low := 0
	high := len(liNew) - 1

	for low <= high {
		newMid := (low + high) / 2

		if liNew[newMid] == target {
			return true
		}
		if liNew[newMid] < target {
			low = newMid + 1
			continue
		}
		high = newMid - 1
	}

	return false
}

func searchMatrixFlatten(matrix [][]int, target int) bool {

	n := len(matrix)
	if n < 0 {
		return false
	}
	m := len(matrix[0])
	if m < 0 {
		return false
	}

	le, ri := 0, n*m-1
	mid := -1
	for le <= ri {
		mid = (le + ri) / 2
		row := mid / m
		col := mid % m
		if matrix[row][col] > target {
			ri = mid - 1
		} else if matrix[row][col] < target {
			le = mid + 1
		} else {
			return true
		}
	}

	return false
}
