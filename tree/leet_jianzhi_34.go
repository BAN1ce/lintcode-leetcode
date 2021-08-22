package tree

// 剑指 Offer 34. 二叉树中和为某一值的路径
// https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
func pathSum(root *TreeNode, target int) [][]int {

	result := make([][]int, 0)
	path := make([]int, 0)

	dfs(root, target, path, &result)
	return result
}

func dfs(root *TreeNode, target int, path []int, result *[][]int) {

	if root == nil {
		return
	}

	path = append(path, root.Val)
	if target == root.Val && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
	}
	if root.Left != nil {
		dfs(root.Left, target-root.Val, path, result)
	}
	if root.Right != nil {
		dfs(root.Right, target-root.Val, path, result)
	}

	path = path[0 : len(path)-1]

}