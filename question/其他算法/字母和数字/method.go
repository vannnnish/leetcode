package 字母和数字

import "fmt"

func findLongestSubarray(array []string) []string {
	// 计算前缀
	var intArr=make([]int,len(array))
	for i,v:=range array{
		if i==0{
			if v[0]>='A'&&v[0]<='z'{
				intArr[i]=1*len(v)
			}else{
				intArr[i]=-1*len(v)
			}
		}else{
			if v[0]>='A'&&v[0]<='z'{
				intArr[i]=intArr[i-1]+1*len(v)
			}else{
				intArr[i]=intArr[i-1]-1*len(v)
			}
		}
	}
	fmt.Println(intArr)
	fmt.Println(intArr[20])
	fmt.Println(intArr[21])
	fmt.Println(intArr[22])
	// 统计最大距离
	distanceMap:=make(map[int]int)
	var max int
	var startIndex,endIndex int
	// 如果数据中没有相同的值，则第一个
	for i,v:=range intArr{
		index,ok:=distanceMap[v]
		if !ok{
			// 如果这个值没存储则存储，
			distanceMap[v]=i
		}else{
			// 如果已经存储过，那么计算此时差值
			tmp:=i-index
			if tmp>max{
				max=tmp
				startIndex=index
				endIndex=i
			}
		}
	}
	return array[startIndex+1:endIndex+1]
}
