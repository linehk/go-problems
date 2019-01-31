package array

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	// map 遍历一次
	// struct{} 是为了减少内存，可以用 bool 代替
	m := make(map[int]struct{})
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = struct{}{}
	}
	return false
}
