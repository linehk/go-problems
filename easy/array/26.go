package array

func removeDuplicates(nums []int) int {
        pmap := make(map[int]int)
        for _, v := range nums {
                if _, ok := pmap[v]; !ok {
                        pmap[v] = len(pmap)
                        nums[pmap[v]] = v
                }
        }
        return len(pmap)
}