package 无重复字符子串

import "strings"

type UniqueQueue struct {
	data [255]rune
	count int
}

func (q *UniqueQueue) Push(a rune) int{
	var resetIndex =-1
	//length:=len(q.data)
	for i:=0;i<q.count;i++{
		if q.data[i]<=0{
			break
		}
		// 出现重复
		if a==q.data[i]{
			resetIndex=i
			break
		}
	}
	if resetIndex!=-1{
		// 数组元素整体向前移动resetIndex位
		for i:=resetIndex;i<q.count-1;i++{
			q.data[i-resetIndex]=q.data[i+1]
		}
		q.data[q.count-resetIndex-1]=a
		q.count=q.count-resetIndex
	}else{
		q.data[q.count]=a
		q.count=q.count+1
	}
	return q.count
}


func Method(s string) int {
	var q =&UniqueQueue{}
	var max int
	for _,v :=range s{
		count:=q.Push(v)
		if count>max{
			max=count
		}
	}
	return max
}

// 别人的方法
/*
核心：只增大不减小的滑动窗口
定义两个游标start和end，遍历字符串，end游标随着遍历无重复增大
若出现重复字符，则两个游标都增大index+1位（窗口大小不变，start游标滑动到重复位置后一位）
这时候由于窗口大小不变，窗口内可能存在重复，恰好遍历从start游标开始，如果没有重复，需要判断i+1与end的关系，超过的话，即i遍历到窗口之外，end再增大
遍历结束，end-start即为结果。
代码
*/
func lengthOfLongestSubstring(s string) int {
	start,end := 0,0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i],string(s[i]))
		if index==-1{
			if i+1>end{
				end=i+1
			}
		}else{
			start+=index+1
			end+=index+1
		}
	}
	return end-start
}
