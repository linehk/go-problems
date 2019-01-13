package string

func firstUniqChar(s string) int {
	var rec [26]int
	for _, v := range s {
		rec[v-'a']++
	}
	for i, v := range s {
		if rec[v-'a'] == 1 {
			return i
		}
	}
	return -1
}
