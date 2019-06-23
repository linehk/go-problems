package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	// 基准条件
	// 无根节点，深度为 0
	if root == nil {
		return 0
	}
	// 左右子树深度的最大值 + 1
	// +1 只是为了符合只有根节点，深度为 1 的定义
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
