package 零和一

import (
	"fmt"
	"testing"
)

func Test_findMaxForm(t *testing.T) {
	strs := []string{"10", "0001", "111001", "1", "0"}
	fmt.Println(findMaxForm(strs, 5, 3))
}
func Test_count(t *testing.T) {
	strs := []string{"10", "0001", "111001", "1", "0"}
	for _, v := range strs {
		fmt.Println(count(v))
	}
}
