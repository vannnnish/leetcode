package 划分字母区间

type Statistic struct {
	firstShowIndex int
	lastShowIndex  int
	count          int
}

func partitionLabels(S string) []int {
	res := make([]int, 0)
	result := make(map[byte]Statistic)
	for i := range S {
		statistic, ok := result[S[i]]
		if !ok {
			statistic.firstShowIndex = i
		}
		statistic.count = statistic.count + 1
		statistic.lastShowIndex = i
		result[S[i]] = statistic
	}
	length := len(S)
	for i := 0; i < length; {
		lastShowIndex := result[S[i]].lastShowIndex
		var count = 1
		if i != lastShowIndex {
			for j := i + 1; ; j++ {
				if j == lastShowIndex {
					count = lastShowIndex - i + 1
					break
				}
				// 如果这段区间内的最后一次出现位置，大于
				if result[S[j]].lastShowIndex >= lastShowIndex {
					lastShowIndex = result[S[j]].lastShowIndex
				}
			}
		}
		i = lastShowIndex + 1
		res = append(res, count)
	}
	return res
}

// LeetCode 官方
func partitionLabels1(s string) (partition []int) {
	lastPos := [26]int{}
	for i, c := range s {
		lastPos[c-'a'] = i
	}
	start, end := 0, 0
	for i, c := range s {
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return
}
