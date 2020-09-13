package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSym(root, root)
}

func isSym(root1, root2 *TreeNode) bool {
	// 终止条件：都为 nil
	if root1 == nil && root2 == nil {
		return true
	}
	// 一个为 nil，另一个不为 nil
	if root1 == nil || root2 == nil {
		return false
	}
	// 值不相同
	if root1.Val != root2.Val {
		return false
	}
	// root1 的左子树和 root2 的右子树是对称的
	// root1 的右子树和 root2 的左子树是对称的
	return isSym(root1.Left, root2.Right) &&
		isSym(root1.Right, root2.Left)
}
