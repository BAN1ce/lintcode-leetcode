package dp

// 剑指 Offer 47. 礼物的最大价值
// https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/
func maxValue(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}
	h, w := len(grid), len(grid[0])
	dp := make([][]int, h)

	for i := 0; i < h; i++ {
		dp[i] = make([]int, w)
	}
	dp[0][0] = grid[0][0]

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if (j-1) > 0 && (i-1) > 0 {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
			if (j - 1) >= 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
			if (i - 1) >= 0 {
				dp[i][j] = max(dp[i][j],dp[i-1][j] + grid[i][j])
			}
		}
	}

	return dp[h-1][w-1]

}
