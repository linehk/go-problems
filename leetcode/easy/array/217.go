package array

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	m := make(map[int]struct{})
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = struct{}{}
	}

	return false
}
