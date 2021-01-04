package 无重叠区间

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	/*
	   	先把区间按照结尾的大小进行增序排序，每次选择结尾最小且和前一个选
	      择的区间不重叠的区间
	*/
	var count int
	length := len(intervals)
	if length == 0 {
		return 0
	}
	//sort.Sort(SortSlice(intervals))
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	// 移除的标记
	compareElement := intervals[0]
	// 进行重叠区间选择
	for i := 1; i < length; i++ {
		// 如果和上一个重叠了，那么肯定是要摘除的
		if intervals[i][0] < compareElement[1] {
			count++
		} else {
			compareElement = intervals[i]
		}
	}
	return count
}

type SortSlice [][]int

func (s SortSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortSlice) Less(i, j int) bool {
	return s[i][1] < s[j][1]
}

func (s SortSlice) Len() int {
	return len(s)
}
