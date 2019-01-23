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
func findInPartiallySortedMatrix2(matrix [][]int, number int) bool {
	if matrix == nil {
		return false
	}
	rows := len(matrix)
	if rows <= 0 {
		return false
	}
	columns := len(matrix[0])
	if columns <= 0 {
		return false
	}
	rowIndex := 0
	columnIndex := columns - 1
	for rowIndex < rows && columnIndex >= 0 {
		if matrix[rowIndex][columnIndex] == number {
			return true
		} else if matrix[rowIndex][columnIndex] > number {
			columnIndex--
		} else {
			rowIndex++
		}
	}
	return false
}
