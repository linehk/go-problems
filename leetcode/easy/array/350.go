package array

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}

	res := []int{}
	for _, v := range nums2 {
		if m[v] > 0 {
			res = append(res, v)
			m[v]--
		}
	}

	return nil
}
