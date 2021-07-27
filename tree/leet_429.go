package tree

type Node struct {
	Val      int
	Children []*Node
}

// N叉树的层序遍历
// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
func levelOrderA(root *Node) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}
	result = append(result, []int{root.Val})
	roots := [][]*Node{root.Children}

	for ; ; {
		level, nextLevel := walkNodeLevel(roots)
		if len(nextLevel) == 0 {
			break
		}
		result = append(result, level)
		roots = nextLevel
	}

	return result
}

func walkNodeLevel(roots [][]*Node) ([]int, [][]*Node) {

	level := make([]int, 0)
	nextLevel := make([][]*Node, 0)
	for _, root := range roots {

		if root == nil {
			continue
		}
		for _, v := range root {
			if v != nil {
				level = append(level, v.Val)
				nextLevel = append(nextLevel, v.Children)
			}
		}
	}
	return level, nextLevel

}
