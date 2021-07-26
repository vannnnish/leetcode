package 汇总

import "sort"

func kthSmallest(mat [][]int, k int) int {
	res := []int{0}
	for _, row := range mat {
		// 对这个
		var newArr []int
		for _, r := range row {
			for _, v := range res {
				newArr = append(newArr, r+v)
			}
		}
		sort.Slice(newArr, func(i, j int) bool {
			return newArr[i] < newArr[j]
		})
		if len(newArr) > k {
			res = newArr[:k]
		} else {
			res = newArr
		}
	}
	return res[len(res)-1]
}
