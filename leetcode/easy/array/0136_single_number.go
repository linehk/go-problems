package array

func singleNumber(nums []int) int {
	res := 0
	for i := range nums {
		// XOR 两个相同的数会得到 0
		// 把所有数都 XOR 一次就可以得到单独的数
		res ^= nums[i]
	}
	return res
}
