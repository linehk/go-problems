package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := new(TreeNode)
	// 根节点的值为中位数
	root.Val = nums[mid]
	// Left 为小的那一半
	root.Left = sortedArrayToBST(nums[:mid])
	// Right 为大的那一半
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
