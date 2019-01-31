package other

func hammingDistance(x int, y int) int {
	x ^= y
	count := 0
	for x != 0 {
		// count += x & 1
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}
