package 青蛙跳台阶

// 递推
func numWaysByRecurrence(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	var sum, a, b = 0, 1, 2
	for i := 3; i <= n; i++ {
		sum = (a + b)%1000000007
		a = b
		b = sum
	}
	return sum
}

// 递归
func numWays(n int) int {
	var dp = make([]int, n+1)
	rel := cal(n, dp)
	return rel % 1000000007
}

func cal(n int, dp []int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	if dp[n] == 0 {
		dp[n] = numWays(n-1) + numWays(n-2)
	}

	return dp[n]
}
