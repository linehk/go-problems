package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// 把链表中全部元素提取到 nums 中
	nums := make([]int, 0)
	for ; head != nil; head = head.Next {
		nums = append(nums, head.Val)
	}
	// 再判断是否是回文
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		if nums[l] != nums[r] {
			return false
		}
	}
	return true
}
