package string

func strStr(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	// 从大的那个字符串中抽取一部分跟 needle 比较
	for i := 0; i <= hlen-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}
	return -1
}
