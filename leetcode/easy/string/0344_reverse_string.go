package string

func reverseString(s []byte) {
	if len(s) < 2 {
		return
	}
	// 用两个指针先中间逼近交换
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
