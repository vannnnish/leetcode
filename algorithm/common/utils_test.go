package common

import (
	"fmt"
	"testing"
)

func TestRandomArrGenerator(t *testing.T) {
	fmt.Println(RandomArrGenerator(10))
	fmt.Println(RandomArrGenerator(20))
}

func TestIsValidSort(t *testing.T) {
	var esc = []int{1, 2, 3, 4, 5, 6, 7}
	var desc = []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println(IsValidSort(desc))
	fmt.Println(IsValidSort(esc))
	fmt.Println(IsValidSort(RandomArrGenerator(5)))
}

func TestDoTimeStatic(t *testing.T) {
	DoTimeStatic(RandomArrGenerator(100000), func(arr []int) {
		for _, _ = range arr {
		}
	})
}
