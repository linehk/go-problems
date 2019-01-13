package string

func reverseString(s string) string {
	ss := []byte(s)
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	return string(ss)
}
