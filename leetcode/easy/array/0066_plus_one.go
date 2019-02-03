package array

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		// 未进位
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// 进位
		digits[i] = 0
	}
	// 最高位进位
	highest := make([]int, len(digits)+1)
	highest[0] = 1
	return highest
}
