package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// fast 指针先走 n 步
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// 避免 fast 走完后链表中只有一个结点
	if fast == nil {
		return head.Next
	}
	// 两个指针一起走直到 fast 指针到尾
	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 现在可以删除 slow 指针了
	slow.Next = slow.Next.Next
	return head
}
