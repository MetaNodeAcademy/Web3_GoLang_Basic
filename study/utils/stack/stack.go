package stack

type Stack struct {
	top  *Node
	size int
}

type Node struct {
	Value interface{}
	Next  *Node
}

func (s *Stack) Push(value interface{}) {
	s.top = &Node{value, s.top}
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	v := s.top.Value
	s.top = s.top.Next
	s.size--
	return v
}

func (s *Stack) Peek() interface{} {
	if s.top == nil {
		return nil
	}
	return s.top.Value
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Size() int {
	return s.size
}
