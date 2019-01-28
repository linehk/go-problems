package other

func missingNumber(nums []int) int {
	xor := 0
	for i, v := range nums {
		xor ^= i ^ v
	}
	return xor ^ len(nums)
}

func missingNumber2(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	n := len(nums)
	return n*(n+1)/2 - sum
}
