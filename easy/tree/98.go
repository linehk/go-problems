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
	return min < root.Val && root.Val < max &&
		isValid(root.Left, min, root.Val) &&
		isValid(root.Right, root.Val, max)
}
