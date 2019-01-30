package array

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
