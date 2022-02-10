package environment

const 수강인원 int = 10

type node struct {
	Value int
}

type Stack struct {
	students []*node
	index    int
	len      int
	limiter  int
}

func NewStack() *Stack {
	return &Stack{students: make([]*node, 수강인원), index: -1, len: 0, limiter: 수강인원}
}

func (s *Stack) Push(n *node) bool {
	if s.len == s.limiter {
		return false
	}
	s.index++
	s.students[s.index] = n
	s.len++
	return true
}

func (s *Stack) Pop() (*node, bool) {
	if s.len == 0 {
		return nil, false
	}
	n := s.students[s.index]
	s.index--
	s.len--
	return n, true
}

func (s *Stack) RenewStack(i int) bool {
	newlimit := s.limiter + i
	newstack := make([]*node, newlimit)
	copy(newstack, s.students)
	s.students = newstack
	s.limiter = newlimit

	return true
}
