package duplicationInArray

import (
	"sort"
)

// 排序法（可得全部重复）
// O(nlgn) O(1)
func duplicationInArray1(nums []int) (int, bool) {
	if len(nums) == 0 {
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
	if len(nums) == 0 {
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
	if len(nums) == 0 {
		return 0, false
	}

	s := make([]int, len(nums))
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
	if len(nums) == 0 {
		return 0, false
	}

	for i := range nums {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i], true
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	return 0, false
}
