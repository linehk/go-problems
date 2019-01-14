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
	queue := list.New()
	queue.PushBack(&NodeInfo{root, 0})
	for queue.Len() != 0 {
		f := queue.Front()
		queue.Remove(f)
		e := f.Value.(*NodeInfo)
		if e.Level == len(res) {
			res = append(res, []int{})
		}
		res[e.Level] = append(res[e.Level], e.Node.Val)
		if e.Node.Left != nil {
			queue.PushBack(&NodeInfo{e.Node.Left, e.Level + 1})
		}
		if e.Node.Right != nil {
			queue.PushBack(&NodeInfo{e.Node.Right, e.Level + 1})
		}
	}
	return res
}
