package tree


// 二叉搜索树结点最小距离
// https://www.lintcode.com/problem/1746/

func minDiffInBST(root *TreeNode) int {

	a := inorderTraversal(root)

	min := abs(a[1] - a[0])

	for i:=1;i<len(a)-1;i++{

		tmp := a[i]- a[i+1]
		if min > abs(tmp) {
			min = abs(tmp)
		}
	}

	return min
}

func abs(a int) int  {
	if a < 0 {
		return  0-a
	}else {
		return a
	}
}