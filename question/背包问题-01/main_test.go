package 背包问题_01

import (
	"fmt"
	"testing"
)

func Test_backPack(t *testing.T){
	arr:=[]int{3,4,8,5}
	size:=10
	fmt.Println(backPack(size,arr))
}
func Test_backPack2(t *testing.T){
	arr:=[]int{3,4,8,5}
	size:=10
	fmt.Println(backPack2(size,arr))
}