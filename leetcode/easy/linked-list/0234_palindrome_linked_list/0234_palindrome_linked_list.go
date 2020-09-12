package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// 把链表中全部元素提取到 nums 中
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	// 再判断是否是回文
	// 回文不是回环，回文是以中间为点，左右两边都相同的
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		if nums[l] != nums[r] {
			return false
		}
	}
	return true
}
