package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 慢指针一次走一步
		fast = fast.Next.Next // 快指针一次走两步
		// 如果它们相遇了则表示有环
		if slow == fast {
			return true
		}
	}
	return false
}
