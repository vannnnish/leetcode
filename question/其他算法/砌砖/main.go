package 砌砖

import "math"

func leastBricks(wall [][]int) int {
	var (
		totalWalls int
		min        int
	)
	totalWalls = len(wall)
	min = math.MaxInt64
	// 记录某个缺口时穿过的砖块
	lengthMap := make(map[int]int)

	// 遍历全部数组，找到所有的截断点
	for i := 0; i < totalWalls; i++ {
		count := len(wall[i])
		var distanceToLeft int
		for j := 0; j < count; j++ {
			distanceToLeft = distanceToLeft + wall[i][j]
			// 统计到左边界距离
			_, ok := lengthMap[distanceToLeft]
			if ok {
				continue
			}
			// 统计每行通过的砖块
			var passBrick int
			for k := 0; k < totalWalls; k++ {
				if k == i {
					continue
				}
				currentCount := len(wall[k])
				if j >= currentCount {
					passBrick++
					continue
				}
				// 统计到具体行次是否通过砖块
				currentBrickLen := 0
				for m := j; m < currentCount; m++ {
					currentCount = currentBrickLen + wall[k][m]
				}
				if currentBrickLen != distanceToLeft {
					passBrick++
				}
			}
		}
	}
	for _, v := range lengthMap {
		if v < min {
			min = v
		}
	}
	return min
}

type mark struct {
	i int
	j int
}
