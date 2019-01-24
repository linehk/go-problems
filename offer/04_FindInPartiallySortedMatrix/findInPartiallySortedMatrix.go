package findInPartiallySortedMatrix

// 暴力法
// O(n2)
func findInPartiallySortedMatrix1(matrix [][]int, number int) bool {
	if matrix == nil {
		return false
	}
	if len(matrix) <= 0 {
		return false
	}
	if len(matrix[0]) <= 0 {
		return false
	}
	for _, rows := range matrix {
		for _, v := range rows {
			if v == number {
				return true
			}
		}
	}
	return false
}

// 一次剔除一列或一行
// O(n)
func findInPartiallySortedMatrix2(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	cols := len(matrix[0])
	if cols == 0 {
		return false
	}

	row := 0
	col := cols - 1
	for row < rows && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
