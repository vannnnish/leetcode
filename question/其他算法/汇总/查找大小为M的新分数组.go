package 汇总

import "fmt"

func findLatestStep(arr []int, m int) int {
	// 用于和原数组大小相同的数组，记录某个点是否被操作过了
	length := len(arr)
	// 记录元素是否被操作过
	recordSlice := make([]int, length)
	// 记录元素位置
	positionRecord := make([][]int, length)
	for i := 0; i < length; i++ {
		positionRecord[i] = []int{-1, -1}
	}
	for _, v := range arr {
		fmt.Println(positionRecord)
		isLeftHasElement, isRightHasElemet := false, false
		recordSlice[v-1] = 0
		// 判断左右两边元素是否被操作过
		if v-2 > 0 {
			if recordSlice[v-2] == 1 {
				isLeftHasElement = true
			}
		}
		// 判断插入点右边是否被操作过
		if v < length {
			if recordSlice[v] == 1 {
				isRightHasElemet = true
			}
		}
		// 如果左右两边都没有元素
		if !isLeftHasElement && !isRightHasElemet {
			positionRecord[v-1] = []int{v - 1, v - 1}
			continue
		}
		// 左右两边都有元素,则是合并区间
		if isLeftHasElement && isRightHasElemet {
			for i := positionRecord[v-2][0]; i <= positionRecord[v][1]; i++ {
				positionRecord[i] = []int{positionRecord[v-2][0], positionRecord[v][1]}
			}
			continue
		}
		// 如果只有左边有值
		if isLeftHasElement {
			for i := positionRecord[v-2][0]; i < v; i++ {
				positionRecord[i] = []int{positionRecord[v-2][0], v}
			}
			continue
		}
		// 如果只有右边有值
		if isRightHasElemet {
			for i := v - 1; i <= positionRecord[v][1]; i++ {
				positionRecord[i] = []int{v - 1, positionRecord[v][1]}
			}
			continue
		}

	}

	return 0
}
