package other

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count++
		// num = num & (num - 1)
		// num = num&num - 1 错误 减号优先级高于与运算符
		// num &= (num - 1)
		// num &= num - 1
		// 类似于 (&=) 的赋值运算符都是先计算右边的, 不需要括号
		num &= num - 1
	}
	return count
}
