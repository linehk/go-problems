package array

func twoSum(nums []int, target int) []int {
	// 用一个 map 来存放索引和判断是否存在某个值
	m := make(map[int]int)
	for i, v := range nums {
		// 得到的 j 为索引
		if j, ok := m[target-v]; ok {
			// j 在 i 前面
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}
