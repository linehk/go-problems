package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	nums := make([]int, 0)
	for ; head != nil; head = head.Next {
		nums = append(nums, head.Val)
	}
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		if nums[l] != nums[r] {
			return false
		}
	}
	return true
}
