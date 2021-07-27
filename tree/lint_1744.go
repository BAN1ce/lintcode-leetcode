package tree

// https://www.lintcode.com/problem/1744/
// 递增顺序查找树
func increasingBST(root *TreeNode) *TreeNode {
	// Write your code here.

	result := inorderTraversal(root)

	tree := &TreeNode{
		Val:   result[0],
		Left:  nil,
		Right: nil,
	}
	next := tree
	for i :=1 ; i < len(result); i++ {
		tmp := &TreeNode{
			Val:   result[i],
			Left:  nil,
			Right: nil,
		}
		next.Right = tmp
		next = tmp
	}

	return tree
}
