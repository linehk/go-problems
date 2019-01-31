package array

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	// 用 map 记录 nums1 中数出现的次数
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}

	res := []int{}
	for _, v := range nums2 {
		// 每遇到一次一个相同的数就往结果数组 append
		if m[v] > 0 {
			res = append(res, v)
			m[v]--
		}
	}
	return res
}
