package tree

// 二叉树的层平均值
// https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/
func averageOfLevels(root *TreeNode) []float64 {

	result := make([]float64, 0)

	if root == nil {
		return result
	}
	roots := []*TreeNode{root}

	result = append(result, float64(root.Val))
	for ; ; {
		level, nextLevel := walkLevel(roots)
		if len(nextLevel) == 0 {
			break
		}
		result = append(result, avg(level))
		roots = nextLevel
	}

	return result

}

func avg(nums []int) float64 {

	sum := 0.0

	for _, v := range nums {
		sum += float64(v)
	}
	return sum / float64(len(nums))
}
