package array

func isValidSudoku(board [][]byte) bool {
	// 平行数组标记是否出现过
	var rowFlag [10][10]bool
	var colFlag [10][10]bool
	var cellFlag [10][10]bool
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				// '9' 的 ascii 码为 57
				// '1' 的为 49
				// '1-9' 的字符减去 '1' 后都在整数 1-9 之间
				// index := board[row][col] - '0' - 1
				index := board[row][col] - '1'
				if rowFlag[row][index] ||
					colFlag[index][col] ||
					cellFlag[row/3*3+col/3][index] {
					return false
				}
				rowFlag[row][index] = true
				colFlag[index][col] = true
				cellFlag[row/3*3+col/3][index] = true
			}
		}
	}
	return true
}
