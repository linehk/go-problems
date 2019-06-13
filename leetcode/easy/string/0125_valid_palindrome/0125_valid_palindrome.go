package string

import (
	"strings"
)

func isPalindrome(s string) bool {
	// 先转换成小写
	s = strings.ToLower(s)
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		// 使 left 指向小写字母
		for left < right && !isChar(s[left]) {
			left++
		}
		// 使 right 指向小写字母
		for left < right && !isChar(s[right]) {
			right--
		}
		// 不相等则表示不是回文
		if s[left] != s[right] {
			return false
		}
	}
	return true
}

func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
