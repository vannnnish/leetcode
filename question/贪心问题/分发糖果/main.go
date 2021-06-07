package 分发糖果

// 已通过测试  https://leetcode-cn.com/problems/candy/
func candy(ratings []int) int {
	var res = make([]int, len(ratings))
	lenOfRatings := len(ratings)
	for i := range ratings {
		res[i] = 1
		if i != 0 {
			if ratings[i] > ratings[i-1] {
				res[i] = res[i-1] + 1
			}
		}
	}

	for i := lenOfRatings - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			res[i-1] = max(res[i]+1, res[i-1])
		}
	}
	var total int
	for i := range res {
		total = total + res[i]
	}
	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
