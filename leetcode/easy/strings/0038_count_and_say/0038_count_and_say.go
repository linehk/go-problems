package string

// 根据上一个数报出下一个数
// 1 1
// 2 11
// 3 21
// 4 1211
// 5 一个1 一个2 两个1 则为 111221
func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	// 循环 n-1 次
	res := []byte{'1'}
	for ; n > 1; n-- {
		cur := []byte{}
		// 遍历上一个结果集
		for i := 0; i < len(res); i++ {
			count := 1
			// 计数
			for i+1 < len(res) && res[i] == res[i+1] {
				count++
				i++
			}
			// byte(count+'0') 把 int 转换为 byte （ASCII码）
			cur = append(cur, byte(count+'0'), res[i])
		}
		// 设置当前结果集为最后结果集
		res = cur
	}
	return string(res)
}
