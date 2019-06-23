package sort_search

// Forward declaration of isBadVersion API.
func isBadVersion(version int) bool {
	return true
}

// 即二分搜索，返回第一个搜索到的
func firstBadVersion(n int) int {
	lo, hi := 1, n
	// 不是 lo <= hi
	for lo < hi {
		mid := lo + (hi-lo)/2
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return hi
}
