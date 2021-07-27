package tree

// 二叉树的最大深度
// https://www.lintcode.com/problem/97/
func maxDepth(root *TreeNode) int {
	// write your code here

	result := make([][]int, 0)

	if root == nil {
		return 0
	}
	result = append(result, []int{root.Val})
	roots := []*TreeNode{root}

	i := 1
	for ; ; {
		_, nextLevel := walkLevel(roots)
		if len(nextLevel) == 0 {
			break
		}
		i++
		roots = nextLevel
	}

	return i
}
