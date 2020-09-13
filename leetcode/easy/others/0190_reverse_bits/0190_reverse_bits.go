package other

// 不断取出 num 最右边的一位然后设置在 res 最右边
// num 不断右移，res 不断左移
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res <<= 1
		res |= num & 1
		num >>= 1
	}
	return res
}
