package array

func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	// 先减一
	n := len(matrix) - 1
	// i 只需遍历一半
	for i := 0; i <= n/2; i++ {
		// i 不断迭代 j 的范围从左右两边收缩
		for j := i; j < n-i; j++ {
			// 先保存左上角
			tmp := matrix[i][j]
			// 把左上角设置为左下角
			matrix[i][j] = matrix[n-j][i]
			// 把左下角设置为右下角
			matrix[n-j][i] = matrix[n-i][n-j]
			// 把右下角设置为右上角
			matrix[n-i][n-j] = matrix[j][n-i]
			// 把右上角设置为左上角
			matrix[j][n-i] = tmp
		}
	}
}
