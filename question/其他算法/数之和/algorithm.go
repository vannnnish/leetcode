/*
@Time : 2020/4/5 15:18
@Author : vannnnish
@File : algorithm
*/

package 数之和

func Method(target int, nums []int) []int {
	if !isValid(nums) {
		panic("nums not valid")
	}
	// 差值map
	diffMap := make(map[int]int)
	for i, v := range nums {
		dif := target - v
		index, ok := diffMap[v]
		if ok {
			return []int{i, index}
		} else {
			diffMap[dif] = i
		}
	}
	return []int{-1, -1}
}

func isValid(nums []int) bool {
	var validMap = map[int]bool{}
	for _, v := range nums {
		if validMap[v] {
			return false
		}
		validMap[v] = true
	}
	return true
}
