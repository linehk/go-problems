package other

func missingNumber1(nums []int) int {
	xor := 0
	for i, v := range nums {
		xor ^= i ^ v
		// 不断地 xor 索引及元素
		// nums = [3, 0, 1]
		// xor1 = 0 ^ (0 ^ 3) = 0000 ^ (0000 ^ 0011) = 0000 ^ 0011 = 0011 = 3
		// xor2 = 3 ^ (1 ^ 0) = 0011 ^ (0001 ^ 0000) = 0011 ^ 0001 = 0010 = 2
		// xor3 = 2 ^ (2 ^ 1) = 0010 ^ (0010 ^ 0001) = 0010 ^ 0011 = 0001 = 1
	}
	// 最后 xor 数组长度
	// xor ^ 3 = 1 ^ 3 = 0001 ^ 0011 = 0010 = 2
	return xor ^ len(nums)
}

func missingNumber2(nums []int) int {
	sum := 0
	// 先取数组和
	for _, v := range nums {
		sum += v
	}
	// nums = [3, 0, 1]
	// sum = 4
	n := len(nums)
	// 用求和公式 = 1+2+3+...+n = n*(n+1)/2
	// 3*(3+1)/2 - 4 = 2
	return n*(n+1)/2 - sum
}
