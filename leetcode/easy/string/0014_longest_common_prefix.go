package string

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	// 遍历最短字符串和切片，拿每个字符进行比较，不等时截止
	shortStr := shortest(strs)
	for cIndex, c := range shortStr {
		for sIndex := range strs {
			// strs[sIndex] 为字符串切片中的字符串
			if strs[sIndex][cIndex] != byte(c) {
				return strs[sIndex][:cIndex]
			}
		}
	}
	return shortStr
}

// 从一个字符串 slice 中返回最短的字符串
func shortest(strs []string) string {
	short := strs[0]
	for _, s := range strs {
		if len(s) < len(short) {
			short = s
		}
	}
	return short
}
