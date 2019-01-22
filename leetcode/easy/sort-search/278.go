package sort_search

func firstBadVersion(n int) int {
	lo, hi := 1, n
	mid := 0
	for lo < hi {
		mid = lo + (hi-lo)/2
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return hi
}
