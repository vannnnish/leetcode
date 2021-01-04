package 根据身高重建队列

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	length := len(people)
	var res = make([][]int, length)
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] < people[j][0]
	})

	for i := 0; i < length; i++ {
		// 统计前面的数据中，是否是满足插入的
		//position := people[i][1]
		var position int
		frontCount := people[i][1]
		var count int
		// 寻找满足个数时的position
		if frontCount > 0 {
			for j := 0; j < length; j++ {
				if len(res[j]) == 0 || res[j][0] >= people[i][0] {
					count++
				}
				if count == frontCount {
					position = j + 1
					break
				}
			}
		}
		// 然后寻找可以插入的position
		for ; position < length; position++ {
			if len(res[position]) == 0 {
				break
			}
		}
		res[position] = people[i]
	}
	return res
}
