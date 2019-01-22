package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	start := head
	start2 := head
	for i := 0; i < n; i++ {
		if start == nil {
			return head
		}
		start = start.Next
	}
	if start == nil {
		return head.Next
	}
	for start.Next != nil {
		start = start.Next
		start2 = start2.Next
	}
	start2.Next = start2.Next.Next
	return head
}
