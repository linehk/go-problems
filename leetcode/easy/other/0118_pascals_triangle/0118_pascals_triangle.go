package other

func generate(numRows int) [][]int {
	res := [][]int{}
	if numRows == 0 {
		return res
	}

	// 初始化第一行
	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}

	for i := 1; i < numRows; i++ {
		// 不断添加行
		res = append(res, getNext(res[i-1]))
	}

	return res
}

// 根据当前行计算生成下一行
func getNext(row []int) []int {
	// 第一个元素为 0，后面的元素来自于 row
	res := make([]int, 1, len(row)+1)
	res = append(res, row...)

	for i := 0; i < len(res)-1; i++ {
		// i 逐个与 i+1 相加
		res[i] += res[i+1]
	}

	return res
}
