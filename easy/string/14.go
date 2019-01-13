package string

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	short := shortest(strs)
	for i, c := range short {
		for j := range strs {
			if strs[j][i] != byte(c) {
				return strs[j][:i]
			}
		}
	}
	return short
}

func shortest(strs []string) string {
	short := strs[0]
	for _, s := range strs {
		if len(short) > len(s) {
			short = s
		}
	}
	return short
}
