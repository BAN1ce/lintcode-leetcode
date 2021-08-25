package tree

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	if  abs(l-r) <= 1{
		return isBalanced(root.Left) && isBalanced(root.Right)
	}else{
		return false
	}

}