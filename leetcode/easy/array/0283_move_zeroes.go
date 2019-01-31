package array

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}
	// slow 指向为 0 的元素
	// fast 遍历一次
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		// 当发现 fast 指向的元素不为 0 时，交换 fast slow 处元素
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
	}
}
