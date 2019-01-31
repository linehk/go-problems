package array

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		// 等于 10 则需要进位
		if digits[i]+1 == 10 {
			digits[i] = 0
			// 最高位也需要进位时
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		} else {
			// 不需要进位之间加
			digits[i]++
			break
		}
	}
	return digits
}
