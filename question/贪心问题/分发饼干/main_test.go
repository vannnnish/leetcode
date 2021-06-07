package 分发饼干

import (
	"fmt"
	"testing"
)

func Test_findContentChildren(t *testing.T) {
	g := []int{1,2,3}
	s := []int{3}
	fmt.Println(findContentChildren2(g, s))
}
