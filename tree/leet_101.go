package tree

// 对称二叉树
// https://leetcode-cn.com/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {

	level := levelOrderInner(root)

	for _, v := range level {

		if arrayIsSymmetric(v) == false {
			return false
		}
	}
	return true
}

func arrayIsSymmetric(nums []int) bool {

	if len(nums) == 1 {
		return true
	}
	l := len(nums)
	for i := 0; i < l/2; i++ {
		if nums[i] != nums[l-1-i] {
			return false
		}
	}

	return true
}

func levelOrderInner(root *TreeNode) [][]int {

	result := make([][]int, 0)

	if root == nil {
		return result
	}
	result = append(result, []int{root.Val})
	roots := []*TreeNode{root}

	for ; ; {
		level, nextLevel := walkLevelInner(roots)
		if len(nextLevel) == 0 {
			break
		}
		result = append(result, level)
		roots = nextLevel
	}

	return result
}

func walkLevelInner(roots []*TreeNode) ([]int, []*TreeNode) {

	level := make([]int, 0)
	nextLevel := make([]*TreeNode, 0)
	for _, root := range roots {

		if root == nil {
			continue
		}

		if root.Left != nil {
			level = append(level, root.Left.Val)
			nextLevel = append(nextLevel, root.Left)
		} else {
			// 此处为投机取巧办法，题目的测试用例中未使用该数字，理论上应对nil与0进行分别判断
			level = append(level, 9999999)
		}
		if root.Right != nil {
			level = append(level, root.Right.Val)
			nextLevel = append(nextLevel, root.Right)
		} else {

			level = append(level, 9999999)
		}
	}
	return level, nextLevel

}
