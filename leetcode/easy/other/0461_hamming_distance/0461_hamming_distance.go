package other

func hammingDistance(x int, y int) int {
	// 先 xor 一下，这时不同位为 1，下面就可以直接计算 x 中 1 的个数
	x ^= y
	count := 0
	for ; x != 0; x >>= 1 {
		if x&1 == 1 {
			count++
		}
	}
	return count
}
