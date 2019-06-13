package array

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	// 快慢指针法
	// 开始时指向相同位置，指向的元素相等时只有 fast 走，不相等时一起走
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			// 把 fast 处不是重复的值复制到 slow
			nums[slow] = nums[fast]
			slow++
		}
	}
	// 索引 0 到 slow 就是结果数组
	return slow
}
