package 汇总

import "math"

/*

i < j 且 min(heights[i], heights[j]) > max(heights[i+1], heights[i+2], ..., heights[j-1])


*/

func canSeePersonsCount(heights []int) []int {
	length := len(heights)
	var ret = make([]int, length)
	for i := 0; i < length; i++ {
		var nextHigh = math.MinInt64
		var tmp int
		for j := i + 1; j < length; j++ {
			// 视野终点
			if heights[i] < heights[j] {
				tmp++
				break
			}
			// 判断这个点是不是能被看到
			if heights[j] > nextHigh {
				nextHigh = heights[j]
			} else {
				continue
			}
			// 统计视野
			// 当前点小于起始点， 且当前点不会被前面的视野挡住
			if heights[i] > heights[j] && heights[j] >= nextHigh {
				tmp++
			}
		}
		ret[i] = tmp
	}
	return ret
}

func canSeePersonsCount1(heights []int) []int {
	n := len(heights)
	ans := make([]int, n)
	stack := []int{math.MaxInt32}
	for i := n - 1; i >= 0; i-- {
		for stack[len(stack)-1] < heights[i] {
			stack = stack[:len(stack)-1]
			ans[i]++
		}
		if len(stack) > 1 { // 还可以再看到一个人
			ans[i]++
		}
		stack = append(stack, heights[i])
	}
	return ans
}
