package tree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	// 64bit min = -1<<63
	// 64bit max = 1<<63-1
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	// 符合 BST 条件：min < root < max
	// 并且左子树也符合，这时根节点值为最大值
	// 并且右子树也符合，这时根节点值为最小值
	return min < root.Val && root.Val < max &&
		isValid(root.Left, min, root.Val) &&
		isValid(root.Right, root.Val, max)
}
