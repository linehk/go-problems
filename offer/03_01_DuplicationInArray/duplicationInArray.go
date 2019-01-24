package duplicationInArray

import (
	"sort"
)

// 排序法（可得全部重复）
// O(nlgn) O(1)
func duplicationInArray1(nums []int) (int, bool) {
	if nums == nil || len(nums) == 0 {
		return 0, false
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i], true
		}
	}

	return 0, false
}

// 哈希表法（可得全部重复）
// O(n) O(n)
func duplicationInArray2(nums []int) (int, bool) {
	if nums == nil || len(nums) == 0 {
		return 0, false
	}

	m := make(map[int]int, len(nums)-1)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v, true
		}
		m[v]++
	}

	return 0, false
}

// 计数法（可得全部重复）
// O(n) O(n)
func duplicationInArray3(nums []int) (int, bool) {
	if nums == nil || len(nums) == 0 {
		return 0, false
	}

	s := make([]int, len(nums)-1)
	for _, v := range nums {
		if s[v] != 0 {
			return v, true
		}
		s[v]++
	}

	return 0, false
}

// 不可得全部重复 需要修改数组
// O(n) O(1)
func duplicationInArray4(nums []int) (int, bool) {
	if nums == nil || len(nums) == 0 {
		return 0, false
	}

	for i, v := range nums {
		for v != i {
			if v == nums[v] {
				return v, true
			}
			nums[i], nums[v] = nums[v], nums[i]
		}
	}

	return 0, false
}
