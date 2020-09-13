package string

import (
	"math"
)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	// 去除空格
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}
	// 设置符号
	if i >= len(str) {
		return 0
	}
	sign := 1
	if str[i] == '+' {
		sign = 1
		i++
	} else if str[i] == '-' {
		sign = -1
		i++
	}
	// 迭代结果
	res := 0
	for ; i < len(str) && '0' <= str[i] && str[i] <= '9'; i++ {
		res = res*10 + (int(str[i]) - '0')
		if sign*res > math.MaxInt32 {
			return math.MaxInt32
		} else if sign*res < math.MinInt32 {
			return math.MinInt32
		}
	}
	return res * sign
}
