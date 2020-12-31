package 收藏清单

import (
	"fmt"
	"sort"
)

/*
	这个实现，效率不行

*/
type LenArr struct {
	OriginIndex int
	sortArr     []string
}

type SortLenArr []LenArr

/*
 // 获取数据集合元素个数
        Len() int
        // 如果 i 索引的数据小于 j 索引的数据，返回 true，且不会调用下面的 Swap()，即数据升序排序。
        Less(i, j int) bool
        // 交换 i 和 j 索引的两个元素的位置
        Swap(i, j int)
*/
func (s SortLenArr) Len() int {
	return len(s)
}
func (s SortLenArr) Less(i, j int) bool {
	if len(s[i].sortArr) > len(s[j].sortArr) {
		return true
	}
	return false
}
func (s SortLenArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type Node struct {
	Node     LenArr
	RootNode *Node
}

func peopleIndexes(favoriteCompanies [][]string) []int {
	rootChildMap := make(map[int]*Node)
	var sortArr SortLenArr
	for i, v := range favoriteCompanies {
		var tmp LenArr
		tmp.OriginIndex = i
		tmp.sortArr = v
		sortArr = append(sortArr, tmp)
	}
	sort.Sort(sortArr)
	// 将数组长度进行分组
	//lastLen := 0
	// 按照数组长度对其进行排序
	for i, strArr := range sortArr {
		// 各个str列表
		if i == 0 {
			rootChildMap[0] = &Node{Node: strArr}
			//lastLen = len(strArr)
		} else {
			// 遍历之前的此slice之前的，
			isFound := false
			for j := i - 1; j >= 0; j-- {
				// 依据题意如果两个数组长度相同，则必不可能是其子串
				if len(strArr.sortArr) == len(sortArr[j].sortArr) {
					continue
				}
				// 如果找到这个几个是前者的子集合，那么说明，也是寻找的的父集合的子集合
				if isSubSet(sortArr[j].sortArr, strArr.sortArr) {
					fmt.Println("j:", j)
					fmt.Println(rootChildMap)
					rootNode, ok := rootChildMap[j]
					if ok {
						rootChildMap[i] = &Node{
							Node:     strArr,
							RootNode: rootNode,
						}
					} else {
						rootChildMap[i] = &Node{
							Node: strArr,
							RootNode: &Node{
								Node:     sortArr[j],
								RootNode: nil,
							},
						}

					}
					isFound = true
					break
				}
			}
			if !isFound {
				rootChildMap[i] = &Node{
					Node:     sortArr[i],
					RootNode: nil,
				}
			}
		}
	}
	var indices = make([]int, 0)
	for _, v := range rootChildMap {
		if v.RootNode == nil {
			indices = append(indices, v.Node.OriginIndex)
		}
	}
	sort.Sort(sort.IntSlice(indices))
	return indices
}

// 判断一个数组是否是另外一个数组子集合
func isSubSet(longerArr, shorterArr []string) bool {
	boolArr := make([]bool, len(shorterArr))
	for i, shorterV := range shorterArr {
		for _, longerV := range longerArr {
			if shorterV == longerV {
				boolArr[i] = true
			}
		}
	}
	var ret = true
	for _, v := range boolArr {
		ret = ret && v
	}
	return ret
}
