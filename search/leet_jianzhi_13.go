package search

// 剑指 Offer 13. 机器人的运动范围
// https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
func movingCount(m int, n int, k int) int {

	board := make([][]int,0)
	for i:=0;i<m;i++{
		board = append(board,make([]int,n))
	}

	return dfs(board,0,0,k)
}

func dfs(board [][]int, i, j, rule int) int {

	h, w := len(board), len(board[0])
	if i >= h || i < 0 || j >= w || j < 0 || !isOk(i, j, rule) || board[i][j] == 1 {
		return 0
	}
	board[i][j] = 1
	return 1 + dfs(board, i+1, j, rule) + dfs(board, i, j+1, rule)
}
func isOk(a, b, rule int) bool {

	return byteSum(a)+byteSum(b) <= rule
}

func byteSum(a int) int {

	sum := 0

	for a != 0 {
		sum += a % 10
		a /= 10
	}
	return sum
}
