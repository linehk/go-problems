package array

func containsDuplicate(nums []int) bool {
        m := make(map[int]struct{})
        for _, v := range nums {
                if _, ok := m[v]; !ok {
                        m[v] = struct{}{}
                } else {
                        return true
                }
        }
        return false
}