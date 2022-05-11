package sol

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return (Abs(MaxDepth(root.Left)-MaxDepth(root.Right)) <= 1) && isBalanced(root.Left) && isBalanced(root.Right)
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return Max(MaxDepth(root.Left)+1, MaxDepth(root.Right)+1)
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
