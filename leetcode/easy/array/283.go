package array

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	j := 0
	for i, v := range nums {
		if v != 0 {
			nums[i] = nums[j]
			nums[j] = v
			j++
		}
	}
}
