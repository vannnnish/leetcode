package 矩阵中的幻方

import (
	"fmt"
	"sort"
	"strconv"
)

func numMagicSquaresInside(grid [][]int) int {
	// 暴力求解
	totalLen := len(grid)
	if totalLen < 3 {
		return 0
	}
	elmLen := len(grid[0])
	if elmLen < 3 {
		return 0
	}

	// 数字出现次数统计
	var showCount = make(map[int]int)
	for i := 9; i <= 9; i++ {
		showCount[i] = 0
	}
	var ret int
	for i := 0; i <= totalLen-3; i++ {
		for j := 0; j <= elmLen-3; j++ {
			// 判断矩阵是不是换3
			stand := grid[i][j] + grid[i][j+1] + grid[i][j+2]
			// 判断行
			if stand != grid[i+1][j]+grid[i+1][j+1]+grid[i+1][j+2] {
				continue
			}
			if stand != grid[i+2][j]+grid[i+2][j+1]+grid[i+2][j+2] {
				continue
			}
			// 判断列
			if stand != grid[i][j]+grid[i+1][j]+grid[i+2][j] {
				continue
			}
			if stand != grid[i][j+1]+grid[i+1][j+1]+grid[i+2][j+1] {
				continue
			}
			if stand != grid[i][j+2]+grid[i+1][j+2]+grid[i+2][j+2] {
				continue
			}
			// 判断对角
			if stand != grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] {
				continue
			}
			if stand != grid[i+2][j]+grid[i+1][j+1]+grid[i][j+2] {
				continue
			}
			if grid[i][j] > 9 || grid[i][j] < 1 {
				continue
			}
			if grid[i][j+1] > 9 || grid[i][j+1] < 1 {
				continue
			}
			if grid[i][j+2] > 9 || grid[i][j+2] < 1 {
				continue
			}
			if grid[i+1][j] > 9 || grid[i+1][j] < 1 {
				continue
			}
			if grid[i+1][j+1] > 9 || grid[i+1][j+1] < 1 {
				continue
			}
			if grid[i+1][j+2] > 9 || grid[i+1][j+2] < 1 {
				continue
			}
			if grid[i+2][j] > 9 || grid[i+2][j] < 1 {
				continue
			}
			if grid[i+2][j+1] > 9 || grid[i+2][j+1] < 1 {
				continue
			}
			if grid[i+2][j+2] > 9 || grid[i+2][j+2] < 1 {
				continue
			}
			// 判断数字是否重复
			showCount[grid[i][j]] = showCount[grid[i][j]] + 1
			showCount[grid[i][j+1]] = showCount[grid[i][j+1]] + 1
			showCount[grid[i][j+2]] = showCount[grid[i][j+2]] + 1
			showCount[grid[i+1][j]] = showCount[grid[i+1][j]] + 1
			showCount[grid[i+1][j+1]] = showCount[grid[i+1][j+1]] + 1
			showCount[grid[i+1][j+2]] = showCount[grid[i+1][j+2]] + 1
			showCount[grid[i+2][j]] = showCount[grid[i+2][j]] + 1
			showCount[grid[i+2][j+1]] = showCount[grid[i+2][j+1]] + 1
			showCount[grid[i+2][j+2]] = showCount[grid[i+2][j+2]] + 1
			if !isLegal(showCount) {
				clear(showCount)
				continue
			}
			ret++
			clear(showCount)
		}
	}
	return ret
}

func toSortedStr(nums []int) {
	var strArr = make([]string, len(nums))
	for i, v := range nums {
		strArr[i] = strconv.Itoa(v)
	}
	fmt.Println(nums)
	sort.Sort(StrArr(strArr))

	fmt.Println(strArr)
}

type StrArr []string

func (arr StrArr) Len() int {
	return len(arr)
}

func (arr StrArr) Less(i, j int) bool {
	str1 := arr[i]
	str2 := arr[j]
	str1Len := len(str1)
	str2Len := len(str2)
	if str1Len == str2Len {
		return str2 < str1
	} else {
		if (str2 + str1) > (str1 + str2) {
			return false
		}
		return true
	}
}

func (arr StrArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 判断是否次数超过
func isLegal(m map[int]int) bool {
	for _, v := range m {
		if v > 1 || v == 0 {
			return false
		}
	}
	return true
}

// 清空次数统计
func clear(m map[int]int) {
	for i := 1; i <= 9; i++ {
		m[i] = 0
	}
}
