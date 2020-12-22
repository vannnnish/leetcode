package 硬笔组合

func Method(n int) int {
	var count = 0
	list := []int{25, 10, 5, 1}
	var valueMap = map[int]int{
		25: 10,
		10: 5,
		5:  1,
	}
	rest := n
	for _, v := range list {
		if v > rest {
			continue
		}
		for rest > 0 {
			if rest < v {
				v = valueMap[v]
			}
			rest = rest - v
			if rest <= 0 {
				count++
				rest = n
				break
			}
		}
	}
	return count
}
