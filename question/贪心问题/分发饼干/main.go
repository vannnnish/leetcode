package 分发饼干

import (
	"sort"
)

// 已通过测试 https://leetcode-cn.com/problems/assign-cookies/
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
		if gi > s[lenOfS-1] || sStart == lenOfS {
			break
		}
		for ; sStart < lenOfS; {
			if s[sStart] >= gi {
				count++
				sStart = sStart + 1
				break
			}
			sStart = sStart + 1
		}
	}
	return count
}

func findContentChildren2(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	lenG := len(g)
	lenS := len(s)
	var ret = 0
	for i, j := 0, 0; i <= lenG-1 && j <= lenS-1; {
		// 满足条件的时候
		if s[j] >= g[i] {
			i++
			j++
			ret++
		} else {
			j++
		}
	}
	return ret
}
