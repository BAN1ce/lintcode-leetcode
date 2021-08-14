package tree

// 剑指 Offer 26. 树的子结构
// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
func isSubStructure(A *TreeNode, B *TreeNode) bool {

	if A == nil || B == nil {
		return false
	}
	if leftRightSame(A, B) == true {
		return true
	}

	if isSubStructure(A.Left, B) == true {
		return true
	}
	if isSubStructure(A.Right, B) == true {
		return true
	}
	return false
}

func leftRightSame(a, b *TreeNode) bool {

	if b == nil {
		return true
	}
	if a == nil {
		return false
	}

	return a.Val == b.Val && leftRightSame(a.Left, b.Left) && leftRightSame(a.Right, b.Right)
}
