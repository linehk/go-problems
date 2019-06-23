package other

func hammingWeight1(num uint32) int {
	count := 0
	for num != 0 {
		count++
		// num = num&num - 1 错误 减号优先级高于与运算符
		// num = num & (num - 1)
		// num &= (num - 1)
		// num &= num - 1
		// 类似于 (&=) 的赋值运算符都是先计算右边的, 不需要括号
		num &= num - 1
	}
	return count
}

// 更直观的解法
func hammingWeight2(num uint32) int {
	count := 0
	for ; num != 0; num >>= 1 {
		if num&1 == 1 {
			count++
		}
	}
	return count
}

// 第三种解法
func hammingWeight3(num uint32) int {
	var count uint32
	for ; num != 0; num >>= 1 {
		count += num & 1
	}
	return int(count)
}
