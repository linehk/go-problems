package math

// 埃拉托斯特尼筛法
func countPrimes(n int) int {
	// 小于 3 的质数数量为 0
	if n < 3 {
		return 0
	}
	// 是否是合数（除了 1 和其本身外具有其他正因数的正整数）
	isComposite := make([]bool, n)
	count := n / 2
	for i := 3; i*i < n; i += 2 {
		if isComposite[i] {
			continue
		}
		for j := i * i; j < n; j += 2 * i {
			if !isComposite[j] {
				count--
				isComposite[j] = true
			}
		}
	}
	return count
}
