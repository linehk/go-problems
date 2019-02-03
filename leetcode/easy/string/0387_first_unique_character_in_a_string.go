package string

func firstUniqChar(s string) int {
	// 建立一个类似 map 的 26 位数组来存放 26 个小写字母
	var m [26]int
	for _, v := range s {
		// v-'a' 是为了转换为 0-25 之间的索引
		m[v-'a']++
	}
	for i, v := range s {
		if m[v-'a'] == 1 {
			return i
		}
	}
	return -1
}
