package tree

// 剑指 Offer 54. 二叉搜索树的第k大节点
// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
func kthLargest(root *TreeNode, k int) int {


	a := midTraversal(root)

	return a[len(a)-1]
}

func midTraversal(root *TreeNode)[]int  {

	a := make([]int, 0)
	if root == nil {
		return a
	}
	if root.Left != nil {
		a = append(a, midTraversal(root.Left)...)
	}
	if root.Right != nil {
		a = append(a, midTraversal(root.Right)...)
	}
	a = append(a, root.Val)
	return a
}
