package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	// 用后面那个指针的值覆盖
	node.Val = node.Next.Val
	// 删除后面那个指针
	node.Next = node.Next.Next
}
