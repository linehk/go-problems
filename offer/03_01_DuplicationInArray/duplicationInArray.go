package duplicationInArray

import (
	"sort"
)

// 排序法（可得全部重复）
// O(nlgn) O(1)
func duplicationInArray1(nums []int) (int, bool) {
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
	m := make(map[int]int, len(nums)-1)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v, true
		} else {
			m[v]++
		}
	}
	return 0, false
}

// 计数法（可得全部重复）
// O(n) O(n)
func duplicationInArray3(nums []int) (int, bool) {
	s := make([]int, len(nums)-1)
	for _, v := range nums {
		if s[v] != 0 {
			return v, true
		} else {
			s[v]++
		}
	}
	return 0, false
}

// 不可得全部重复 需要修改数组
// O(n) O(1)
func duplicationInArray4(nums []int) (int, bool) {
	for i, v := range nums {
		if v != i {
			if v == nums[v] {
				return v, true
			} else {
				nums[i], nums[v] = nums[v], nums[i]
			}
		}
	}
	return 0, false
}
