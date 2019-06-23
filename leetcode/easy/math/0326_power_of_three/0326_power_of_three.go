package math

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	// 当 n%3 != 0 时停止循环
	for n%3 == 0 {
		n /= 3
	}
	// n == 1 则表示是
	// n == 2 则表示不是
	return n == 1
}
