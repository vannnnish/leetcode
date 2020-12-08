package skiplist

import (
	"math/rand"
)

const (
	maxLevel int     = 16
	p        float32 = 0.25
)

func randomLevel() int {
	level := 1
	for rand.Float32() < p && level < maxLevel {
		level++
	}
	return level
}

type Node struct {
	Score   float64
	Value   interface{}
	forward []*Node
}

func newElement(score float64, value interface{}, level int) *Node {
	return &Node{
		Score:   score,
		Value:   value,
		forward: make([]*Node, level),
	}
}

type SkipList struct {
	header *Node // 虚拟头结点
	len    int   // 跳跃表的长度，不包含虚拟头结点
	level  int   // 当前跳跃表的层数，不包含虚拟头结点
}

func New() *SkipList {
	return &SkipList{
		header: &Node{
			forward: make([]*Node, maxLevel),
		},
	}
}

// 返回长度
func (s *SkipList) Size() int {
	return s.len
}

// 返回跳跃表
func (s *SkipList) Front() *Node {
	return s.header.forward[0]
}

// 返回e之后的元素
func (e *Node) Next() *Node {
	if e != nil {
		return e.forward[0]
	}
	return nil
}

// 查找
func (s *SkipList) Search(score float64) (element *Node, ok bool) {
	x := s.header
	for i := s.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil && x.Score == score {
		return x, true
	}
	return nil, false
}

// 插入
func (s *SkipList) Insert(score float64, value interface{}) *Node {
	update := make([]*Node, maxLevel)
	x := s.header
	// 先查找插入点
	for i := s.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	// 如果Score已经存在，则替换新的值
	if x != nil && x.Score == score {
		x.Value = value
		return x
	}
	// 生成插入节点的level
	level := randomLevel()
	if level > s.level {
		level = s.level + 1
		update[s.level] = s.header
		s.level = level
	}
	e := newElement(score, value, level)
	for i := 0; i < level; i++ {
		e.forward[i] = update[i].forward[i]
		update[i].forward[i] = e
	}
	s.len++
	return e
}

func (sl *SkipList) Delete(score float64) *Node {
	update := make([]*Node, maxLevel)
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]

	if x != nil && x.Score == score {
		for i := 0; i < sl.level; i++ {
			if update[i].forward[i] != x {
				return nil
			}
			update[i].forward[i] = x.forward[i]
		}
		sl.len--
	}
	return x
}