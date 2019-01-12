package array

func removeDuplicates(nums []int) int {
        m := make(map[int]int)
        for _, v := range nums {
                if _, ok := m[v]; !ok {
                        m[v] = len(m)
                        nums[m[v]] = v
                }
        }
        return len(m)
}