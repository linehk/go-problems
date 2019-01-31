package string

func reverseString(s []byte) {
	if len(s) < 2 {
		return
	}
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
