package 汇总

import (
	"sort"
)

func bestTeamScore(scores []int, ages []int) int {
	l := len(scores)
	max := 0
	arr := make([][2]int, l)
	for i, v := range ages {
		arr[i] = [2]int{v, scores[i]}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0] || (arr[i][0] == arr[j][0] && arr[i][1] < arr[j][1])
	})
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j][1] <= arr[i][1] {
				if dp[i] < dp[j] {
					dp[i] = dp[j]
				}
			}
		}
		dp[i] += arr[i][1]
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
