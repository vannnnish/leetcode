package common

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomArrGenerator(num int) []int {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	ints := make([]int, num)
	for i := range ints {
		ints[i] = r.Intn(num)
	}
	return ints
}

func InOrderArrGenerator(num int) []int {
	ints := make([]int, num)
	for i := num - 1; i > 0; i-- {
		ints[i] = i
	}
	return ints
}

// 重复度比较高的数据
func RepeatNumberArrGenerator(num int, max int) []int {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	ints := make([]int, num)
	for i := range ints {
		ints[i] = r.Intn(max)
	}
	return ints
}

func IsValidSort(nums []int) bool {
	length := len(nums)
	isDesc, isEsc := true, true
	// 从大到小
	for i := 0; i < length-1; i++ {
		if nums[i] < nums[i+1] {
			isDesc = false
			break
		}
	}
	// 从小到大
	for i := 0; i < length-1; i++ {
		if nums[i] > nums[i+1] {
			isEsc = false
			break
		}
	}
	return isDesc || isEsc
}

func DoTimeStatic(arr []int, f func(arr []int)) {
	length := len(arr)
	now := time.Now()
	f(arr)
	fmt.Printf("数据量:%d, 排序用时:%0.4fms\n", length, time.Since(now).Seconds()*1000)

}
