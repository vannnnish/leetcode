package 最小路径和

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	// 定义数据
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 初始化值
	for i := 0; i < m; i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	for j := 0; j < n; j++ {
		if j == 0 {
			dp[0][j] = grid[0][j]
		} else {
			dp[0][j] = dp[0][j-1] + grid[0][j]
		}
	}
	// 递推求解
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] > dp[i][j-1] {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}
