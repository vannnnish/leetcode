package 零和一

import (
	"strings"
)

/*
给你一个二进制字符串数组 strs 和两个整数 m 和 n 。

请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。

如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。

*/

func findMaxForm(strs []string, m int, n int) int {
	length := len(strs)
	dp := make([][][]int, length+1)
	for i := 0; i <= length; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}
	// 递归
	for i := 1; i <= length; i++ {
		zeroCount, oneCount := count(strs[i-1])
		for j := 1; j <= m; j++ {
			for k := 1; k <= n; k++ {
				if j >= zeroCount && k >= oneCount {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-zeroCount][k-oneCount]+1)
				} else {
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}

	return dp[length][m][n]
}

// 统计字符串1和0数量
func count(str string) (int, int) {
	var zeroCount, oneCount int
	strings.Map(func(r rune) rune {
		if r == '0' {
			zeroCount++
			return r
		}
		if r == '1' {
			oneCount++
			return r
		}
		return r
	}, str)
	return zeroCount, oneCount
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 优化成二维
func findMaxForm2(strs []string, m int, n int) int {
	statistic := func(str string) (int, int) {
		var zero, one = 0, 0
		for _, char := range str {
			if char == '0' {
				zero++
			} else {
				one++
			}
		}
		return zero, one
	}
	dp := make([][]int, m+1)
	//初始化
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	//进行动态规划
	for _, str := range strs {
		zero, one := statistic(str)
		//当zero或者one比m或者n大时，这个字符串就不用判断了。
		if zero > m || one > n {
			continue
		}
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
			}
		}
	}
	return dp[m][n]
}
