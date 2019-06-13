package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode  // 前一个结点
	cur := head        // 当前结点
	var next *ListNode // 下一个结点
	for cur != nil {
		next = cur.Next // 初始化 next 的位置
		cur.Next = pre  // 设置当前结点的下一个结点为前一个结点
		pre = cur       // pre 前进
		cur = next      // cur 前进
	}
	return pre
}
