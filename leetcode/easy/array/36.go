package array

func isValidSudoku(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		if !isValidSudokuRow(board, row) {
			return false
		}
	}

	for col := 0; col < 9; col++ {
		if !isValidSudokuCol(board, col) {
			return false
		}
	}

	for pod := 0; pod < 9; pod++ {
		if !isValidSudokuPod(board, pod) {
			return false
		}
	}

	return true
}

func isValidSudokuRow(board [][]byte, row int) bool {
	var isExist [10]bool
	for col := 0; col < 9; col++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if isExist[n] {
			return false
		}
		isExist[n] = true
	}
	return true
}

func isValidSudokuCol(board [][]byte, col int) bool {
	var isExist [10]bool
	for row := 0; row < 9; row++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if isExist[n] {
			return false
		}
		isExist[n] = true
	}
	return true
}

func isValidSudokuPod(board [][]byte, pod int) bool {
	var isExist [10]bool
	row := (pod / 3) * 3
	col := (pod % 3) * 3
	for podRow := 0; podRow < 3; podRow++ {
		for podCol := 0; podCol < 3; podCol++ {
			n := convertToNumber(board[row+podRow][col+podCol])
			if n < 0 {
				continue
			}
			if isExist[n] {
				return false
			}
			isExist[n] = true
		}
	}
	return true
}

func convertToNumber(b byte) int {
	if b == '.' {
		return -1
	}
	return int(b - '0')
}
