package graph

import (
	"fmt"
	"testing"
)

func TestMatrix(t *testing.T) {
	m := &MatrixRepeat{}
	m.ReadFromFile("./g.txt")
	fmt.Println(m.LinkedVertex(1))
	fmt.Println(m)
}
