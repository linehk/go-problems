package array

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	// 取余操作是为了在 k > len(nums) 的情况也能正常运行
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

// 翻转从索引 i 到 j 的数
func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
