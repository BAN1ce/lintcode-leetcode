package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	a := make([]int, 0)
	if root == nil {
		return a
	}
	if root.Left != nil {
		a = append(a, inorderTraversal(root.Left)...)
	}
	a = append(a, root.Val)
	if root.Right != nil {
		a = append(a, inorderTraversal(root.Right)...)
	}
	return a
}

func TMain(root *TreeNode) interface{} {
	return levelOrder(root)
}
