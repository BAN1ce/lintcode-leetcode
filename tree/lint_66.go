package tree

// https://www.lintcode.com/problem/66/
// 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {

	a := make([]int, 0)
	if root == nil {
		return a
	}
	a = append(a, root.Val)
	if root.Left != nil {
		a = append(a, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		a = append(a, preorderTraversal(root.Right)...)
	}
	return a
}
