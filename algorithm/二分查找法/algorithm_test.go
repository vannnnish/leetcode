package 二分查找法

import (
	"fmt"
	"testing"
)

func TestBinaryFind(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 39}
	fmt.Println(BinarySearch2(input, 39 ))
}
