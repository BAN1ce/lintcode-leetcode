package tree

// 剑指 Offer 27. 二叉树的镜像
// https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
func mirrorTree(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)

	return root
}
