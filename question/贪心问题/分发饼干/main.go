package 分发饼干

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	var count int
	sort.Ints(g)
	sort.Ints(s)
	lenOfS := len(s)
	if lenOfS == 0 {
		return 0
	}
	var sStart int
	for _, gi := range g {
		// 如果胃口比最大的都还大，那么直接跳过了
		if gi > s[lenOfS-1] || sStart==lenOfS{
			break
		}
		for ; sStart < lenOfS;  {
			if s[sStart] >= gi {
				count++
				sStart=sStart+1
				break
			}
			sStart=sStart+1
		}
	}
	return count
}
