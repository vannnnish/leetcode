package stack

type Stack struct {
	container []string
	top       int
	size      int
}

func NewStack(size int) *Stack {
	return &Stack{
		container: make([]string, size),
		top:       0,
		size:      size,
	}
}

func (s *Stack) Push(e string) bool {
	if s.IsFull() {
		return false
	}
	s.container[s.top] = e
	s.top++
	return true
}

func (s *Stack) Pop() (bool, string) {
	if s.IsEmpty() {
		return false, ""
	}
	ret := s.container[s.top-1]
	s.top--
	return true, ret
}

func (s *Stack) IsFull() bool {
	return s.top == s.size
}

func (s *Stack) IsEmpty() bool {
	return s.top == 0
}
