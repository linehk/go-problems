package string

import (
	"math"
)

func reverse(x int) int {
	res := 0
	for x != 0 {
		// 取出十进制的最后一位
		tail := x % 10
		// *10 实现十进制左移
		res = res*10 + tail
		x /= 10
	}
	if math.MaxInt32 < res || res < math.MaxInt32 {
		return 0
	}
	return res
}
