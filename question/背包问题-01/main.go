package 背包问题_01

/*
有 N 件物品和一个容量是 V 的背包。每件物品只能使用一次。

第 i 件物品的体积是 vi，价值是 wi。

求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。
输出最大价值。

输入格式
第一行两个整数，N，V，用空格隔开，分别表示物品数量和背包容积。

接下来有 N 行，每行两个整数 vi,wi，用空格隔开，分别表示第 i 件物品的体积和价值。

输出格式
输出一个整数，表示最大价值。
*/

func backPack(m int, A []int) int {
	// write your code here
	length := len(A)
	dp := make([][]int, length+1)
	for i := 0; i <= length; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= length; i++ {
		for j := 1; j <= m; j++ {
			if j >= A[i-1] {
				//
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-A[i-1]]+A[i-1])
			} else {
				// 不选
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[length][m]
}

func backPack2(m int, A []int) int {
	// write your code here
	length := len(A)
	dp := make([]int, m+1)

	for i := 1; i <= length; i++ {
		for j := m; j >= A[i-1]; j-- {
			dp[j] = max(dp[j], dp[j-A[i-1]]+A[i-1])

		}
	}
	return dp[m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
