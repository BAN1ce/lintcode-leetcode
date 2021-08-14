package tree

// 剑指 Offer 28. 对称的二叉树
// https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
func isSymmetric(root *TreeNode) bool {

	return treeSymmeric(root, root)
}

func treeSymmeric(a, b *TreeNode) bool {

	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && treeSymmeric(a.Left, b.Right) && treeSymmeric(a.Right,b.Left)
}
