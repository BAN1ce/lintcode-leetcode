package tree

// 226 翻转二叉树
// https://leetcode-cn.com/problems/invert-binary-tree/
func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root

}
