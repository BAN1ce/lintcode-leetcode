package tree


// 剑指 Offer 32 - III. 从上到下打印二叉树 III
// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
func levelOrderv3(root *TreeNode) [][]int {

	result := make([][]int, 0)

	if root == nil {
		return result
	}
	result = append(result, []int{root.Val})
	roots := []*TreeNode{root}

	i := 1
	for ; ; {
		level, nextLevel := walkLevel1(roots)
		if len(nextLevel) == 0 {
			break
		}
		if i%2 == 1 {

			result = append(result, reverse(level))
		} else {
			result = append(result, level)
		}
		i++
		roots = nextLevel
	}

	return result
}

func walkLevel1(roots []*TreeNode) ([]int, []*TreeNode) {

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

func reverse(a []int) []int {

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
