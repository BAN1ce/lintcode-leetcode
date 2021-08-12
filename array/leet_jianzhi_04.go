package array

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
