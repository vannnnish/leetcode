/*
@Time : 2020/4/5 16:08
@Author : vannnnish
@File : algorithm_test
*/

package 二进制手表

import (
	"fmt"
	"testing"
)

func TestMethod(t *testing.T) {
	fmt.Println(Method(1))
}

func TestConcertToBin(t *testing.T){
	fmt.Println(convertToBin(0))
	fmt.Println(convertToBin(1))
	fmt.Println(convertToBin(3))
	fmt.Println(convertToBin(2))
	fmt.Println(convertToBin(9))
}