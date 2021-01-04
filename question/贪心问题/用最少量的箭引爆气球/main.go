package 用最少量的箭引爆气球

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	var (
		lastCompare []int
		count       = 1
	)
	length := len(points)
	if length == 0 {
		count = 0
	} else {
		lastCompare = points[0]
	}
	for i := 1; i < length; i++ {
		if points[i][0] > lastCompare[1] {
			lastCompare = points[i]
			count++
		}
	}
	return count
}
