package tree

// 二叉树的层次遍历
// https://www.lintcode.com/problem/69/
func levelOrder(root *TreeNode) [][]int {

	result := make([][]int, 0)

	if root == nil {
		return result
	}
	result = append(result, []int{root.Val})
	roots := []*TreeNode{root}

	for ; ; {
		level, nextLevel := walkLevel(roots)
		if len(nextLevel) == 0 {
			break
		}
		result = append(result, level)
		roots = nextLevel
	}

	return result
}

func walkLevel(roots []*TreeNode) ([]int, []*TreeNode) {

	level := make([]int, 0)
	nextLevel := make([]*TreeNode, 0)
	for _, root := range roots {

		if root == nil {
			continue
		}

		if root.Left != nil {
			level = append(level, root.Left.Val)
			nextLevel = append(nextLevel, root.Left)
		}
		if root.Right != nil {
			level = append(level, root.Right.Val)
			nextLevel = append(nextLevel, root.Right)
		}

	}
	return level, nextLevel

}
