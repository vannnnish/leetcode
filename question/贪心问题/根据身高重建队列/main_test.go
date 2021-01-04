package 根据身高重建队列

import (
	"fmt"
	"testing"
)

func Test_reconstructQueue(t *testing.T) {
	input := [][]int{
		{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2},
	}
	fmt.Println(reconstructQueue(input))
}
