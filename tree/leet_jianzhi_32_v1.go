package tree

// 剑指 Offer 32 - I. 从上到下打印二叉树
// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/

func levelOrder32(root *TreeNode) []int {

	result := make([]int, 0)

	if root == nil {
		return result
	}
	result = append(result, root.Val)
	roots := []*TreeNode{root}

	for ; ; {
		level, nextLevel := walkLevel32(roots)
		if len(nextLevel) == 0 {
			break
		}
		result = append(result, level...)
		roots = nextLevel
	}

	return result
}

func walkLevel32(roots []*TreeNode) ([]int, []*TreeNode) {

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
