package math

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
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
