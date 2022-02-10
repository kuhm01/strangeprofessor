package environment

type node struct {
	Value int
}

type stack struct {
	students []*node
	index    int
	len      int
	limiter  int
}

func NewStack() *stack {
	return &stack{students: make([]*node, 30), index: -1, len: 0, limiter: 30}
}

func (s *stack) Push(n *node) bool {
	if s.len == s.limiter {
		return false
	}
	s.index++
	s.students[s.index] = n
	s.len++
	return true
}

func (s *stack) Pop() (*node, bool) {
	if s.len == 0 {
		return nil, false
	}
	n := s.students[s.index]
	s.index--
	s.len--
	return n, true
}

func (s *stack) RenewStack(i int) bool {
	newlimit := s.limiter + i
	newstack := make([]*node, newlimit)
	copy(newstack, s.students)
	s.students = newstack
	s.limiter = newlimit

	return true
}
