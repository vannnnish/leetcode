/*
@Time : 2020/4/5 15:18
@Author : vannnnish
@File : algorithm
*/

package 打家劫舍

import "fmt"

func Method(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return nums[0]
	}
	dp := make([]int, length)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[length-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 空间优化
func MethodOpt(nums []int) int {
	length := len(nums)
	var res int
	var beforeRes int
	res = nums[0]
	for i := 1; i < length; i++ {
		tmp := max(res, beforeRes+nums[i])
		beforeRes = res
		res = tmp
	}
	return res
}

// 打家劫舍2.0 如果是首尾相连的情况
func rob(nums []int) int {

	length := len(nums)
	var leftRes, rightRes = nums[0], nums[length-1]
	// 从左到右
	var beforeLeft, beforeRight int
	for i := 1; i < length-1; i++ {
		tmp := max(leftRes, beforeLeft+nums[i])
		beforeLeft = leftRes
		leftRes = tmp
	}
	// 从右到左
	for i := length - 2; i > 0; i-- {
		tmp := max(rightRes, beforeRight+nums[i])
		beforeRight = rightRes
		rightRes = tmp
	}
	if rightRes > leftRes {
		return rightRes
	}
	return leftRes
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 打家劫舍 3.0
func rob3(root *TreeNode) int {
	 return 0
}

// 先序遍历
func preTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preTraverse(root.Left)
	preTraverse(root.Right)
}

func traverseLeft(root *TreeNode){

}
