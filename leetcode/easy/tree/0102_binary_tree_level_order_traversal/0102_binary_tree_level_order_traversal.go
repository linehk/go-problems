package tree

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeInfo struct {
	Node  *TreeNode
	Level int
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	// 层序遍历用 FIFO 队列来做边缘集
	queue := list.New()
	// 在队列中加入根节点，根节点的 Level = 0
	queue.PushBack(&NodeInfo{root, 0})
	for queue.Len() != 0 {
		// 出队
		f := queue.Front()
		queue.Remove(f)
		e := f.Value.(*NodeInfo)
		// 新建一层
		if e.Level == len(res) {
			res = append(res, []int{})
		}
		// 往该层添加结点值
		// res[e.Level] 不会 panic，因为会新建层
		res[e.Level] = append(res[e.Level], e.Node.Val)
		// 如果左子节点不为空，则也插入队列，Level + 1
		if e.Node.Left != nil {
			queue.PushBack(&NodeInfo{e.Node.Left, e.Level + 1})
		}
		// 如果右子节点不为空，则也插入队列，Level + 1
		if e.Node.Right != nil {
			queue.PushBack(&NodeInfo{e.Node.Right, e.Level + 1})
		}
	}
	return res
}
