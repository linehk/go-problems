package string

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	src := []byte{'1'}
	for i := 2; i <= n; i++ {
		src = cas(src)
	}
	return string(src)
}

func cas(src []byte) []byte {
	des := make([]byte, 0)
	x, y := 0, 1
	for x < len(src) {
		for y < len(src) && src[x] == src[y] {
			y++
		}
		des = append(des, byte(y-x+'0'), src[x])
		x = y
	}
	return des
}
