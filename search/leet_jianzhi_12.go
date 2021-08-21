package search

// 剑指 Offer 12. 矩阵中的路径
// https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
func exist(board [][]byte, word string) bool {

	if len(board) == 0 {

		return false
	}
	h, w := len(board), len(board[0])

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if dfs1(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs1(board [][]byte, i, j int, word string, index int) bool {

	h, w := len(board), len(board[0])
	if i >= h || i < 0 || j >= w || j < 0 || board[i][j] != word[index] {
		return false
	}
	if index == len(word)-1 {
		return true
	}
	board[i][j] = '0'

	result := dfs1(board, i+1, j, word, index+1) || dfs1(board, i-1, j, word, index+1) || dfs1(board, i, j+1, word, index+1) || dfs1(board, i, j-1, word, index+1)
	board[i][j] = word[index]

	return result
}
