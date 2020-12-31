package 和为奇数的子数组个数

import "fmt"

/*
给你一个整数数组 arr 。请你返回和为 奇数 的子数组数目。

由于答案可能会很大，请你将结果对 10^9 + 7 取余后返回。
*/

func numOfSubarrays(arr []int) int {
	// 统计奇数个数
	var oddCount, evenCount int
	for _, v := range arr {
		if v%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	fmt.Println("even:", evenCount)
	fmt.Println("odd:", oddCount)
	// 奇数排列方式
	oddWay := 0
	for i := 1; i <= oddCount; i++ {
		if i%2 == 0 {
			continue
		}
		oddWay += combination(oddCount, i)
	}

	// 偶数全排列
	return oddWay * factorial(evenCount)
}

// 求组合  n里面取m
func combination(n, m int) int {
	//var denominator, numerator = n, m
	// 求分子
	for i := 1; i < m; i++ {
		n *= n - i
		m *= i
	}
	// 求分母
	return n / m
}

// 求阶乘
func factorial(n int) int {
	var ret int
	if n < 0 {
		return 0
	}
	for i := 1; i < n; i++ {
		ret *= i
	}
	return ret
}
