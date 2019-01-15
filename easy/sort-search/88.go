package sort_search

func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, len(nums1))
	copy(temp, nums1)
	i, j := 0, 0
	for k := 0; k < len(nums1); k++ {
		if i >= m {
			nums1[k] = nums2[j]
			j++
		} else if j >= n {
			nums1[k] = temp[i]
			i++
		} else if temp[i] < nums2[j] {
			nums1[k] = temp[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
	}
}
