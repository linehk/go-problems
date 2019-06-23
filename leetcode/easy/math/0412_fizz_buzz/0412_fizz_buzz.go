package math

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := range res {
		x := i + 1
		switch {
		// 同时是 3 和 5 的倍数，即 15
		case x%15 == 0:
			res[i] = "FizzBuzz"
		case x%5 == 0:
			res[i] = "Buzz"
		case x%3 == 0:
			res[i] = "Fizz"
		default:
			// 正常数字
			res[i] = strconv.Itoa(x)
		}
	}
	return res
}
