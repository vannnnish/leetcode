package 最小操作次数使元素值相等

import "math"

/*

给定一个长度为 n 的 非空 整数数组，每次操作将会使 n - 1 个元素增加 1。找出让数组所有元素相等的最小操作次数。



示例：

输入：
[1,2,3]
输出：
3
解释：
只需要3次操作（注意每次操作会增加两个元素的值）：
[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]

*/
// 会超时
func minMoves(nums []int) int {
	maxIndex := 0
	maxValue := nums[0]
	length := len(nums)
	opTimes := 0
	for {
		if isDone(nums) {
			break
		}
		// 找到最大值的序号
		for i := 1; i < length; i++ {
			if nums[i] > maxValue {
				maxValue = nums[i]
				maxIndex = i
			}
		}
		// 进行元素加法操作
		for i := range nums {
			if i == maxIndex {
				continue
			}
			nums[i] = nums[i] + 1
		}
		opTimes++

	}
	return opTimes
}

func isDone(nums []int) bool {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		if nums[i] != nums[i+1] {
			return false
		}
	}
	return true
}

func mathWay(nums []int) int {
	length := len(nums)
	min := math.MaxInt64
	moves := 0
	for i := range nums {
		moves += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
	}
	return moves - length*min
}
