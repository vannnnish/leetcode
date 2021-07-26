package 汇总

func maxHappyGroups(batchSize int, groups []int) int {
	// 暴力求解
	var newResult []int
	result := 0
	for _, v := range groups {
		if v%batchSize == 0 {
			result++
		} else {
			newResult = append(newResult, v%batchSize)
		}
	}

	// 用于记录对应下标的元素已经使用过了
	recordMap := make(map[int]bool)
	length := len(newResult)
	// 两组配合能满足的
	for i, v := range newResult {
		// 需要在余下的元素中找到补充的元素
		if _, ok := recordMap[i]; ok {
			continue
		}
		needCompant := v % batchSize
		for j := i + 1; j < length; j++ {
			if _, ok := recordMap[j]; ok {
				continue
			}
			if newResult[j]+needCompant == batchSize {
				// 找到恰好能一次满足的
				recordMap[i] = true
				recordMap[j] = true
				result++
			}
		}
	}
	// 对结果进行统计
	var afterTwoResult []int
	for i := range newResult {
		if _, ok := recordMap[i]; !ok {
			afterTwoResult = append(afterTwoResult, newResult[i])
		}
	}
	return 1
}
