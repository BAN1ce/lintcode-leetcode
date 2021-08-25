package tree

func maxDepth(root *TreeNode) int {

	max := 0
	dps2(root,0,&max)
	return max
}

func dps2(root *TreeNode, value int, max *int) int {
	if root != nil {
		value += 1
		if value > *max {
			*max = value
		}
		return dps2(root.Left, value, max) + dps2(root.Right, value, max)
	} else {
		return -1
	}
}
