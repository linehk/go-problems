package string

// ASCII版本
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var m [26]int32
	for _, v := range s {
		m[v-'a']++
	}
	for _, v := range t {
		if m[v-'a'] == 0 {
			return false
		}
		m[v-'a']--
	}
	return true
}

// UTF版本
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 先建立一个 map 用来计算字符出现次数
	m := make(map[rune]int, len(s))
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		// 当在 s 出现的在 t 却没有出现
		if m[v] == 0 {
			return false
		}
		// 在 t 出现了就对应减一
		m[v]--
	}
	return true
}
