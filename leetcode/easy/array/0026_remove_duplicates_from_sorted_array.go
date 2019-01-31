package array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 快慢指针法
	// 开始时指向相同位置，指向的元素相等时只有 fast 走，不相等时一起走
	slow := 0
	// 第一个元素和自己比较肯定相等，fast 就不用从 0 开始了
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			// 把 fast 处不是重复的值复制到 slow
			nums[slow] = nums[fast]
		}
	}
	// 索引 0 到 slow 就是结果数组
	return slow + 1
}
