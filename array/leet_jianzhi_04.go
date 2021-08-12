package array

// 剑指 Offer 04. 二维数组中的查找
// https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
func findNumberIn2DArray(matrix [][]int, target int) bool {

	if len(matrix) == 0 {
		return false
	}

	for i:=0;i<len(matrix[0]);i++{
		for j := 0 ;j < len(matrix);j++{
			if matrix[j][i] > target {
				break
			}
			if matrix[j][i] == target {
				return true
			}
		}
	}
	return false
}
