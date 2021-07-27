package tree


// https://www.lintcode.com/problem/68/
// 二叉树的后序遍历
func postorderTraversal (root *TreeNode) []int {
	// write your code here

	a := make([]int, 0)
	if root == nil {
		return a
	}
	if root.Left != nil {
		a = append(a, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		a = append(a, postorderTraversal(root.Right)...)
	}
	a = append(a, root.Val)
	return a
}

