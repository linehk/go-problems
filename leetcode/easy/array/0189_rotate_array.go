package array

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	// 取余操作是为了在 k > len(nums) 的情况也能正常运行
	// 左移 1. 翻转 0 到 k 2. 翻转剩下元素 3. 翻转全部
	// 右移 先进行第三步
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
