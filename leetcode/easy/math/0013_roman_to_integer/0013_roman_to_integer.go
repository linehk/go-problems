package math

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0
	last := 0
	// 从字符串后面开始
	for i := len(s) - 1; i >= 0; i-- {
		temp := m[s[i]]
		// 默认为正数，当 temp < last 时转换符号
		sign := 1
		if temp < last {
			sign = -1
		}
		res += sign * temp
		// last 递进
		last = temp
	}
	return res
}
