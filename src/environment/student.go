package environment

import "fmt"

const 수강인원 int = 10

type Node struct {
	Value int
}

type Stack struct {
	students []*Node
	index    int
	len      int
	limiter  int
}

func NewStack() *Stack {
	return &Stack{students: make([]*Node, 수강인원), index: -1, len: 0, limiter: 수강인원}
}

func (s *Stack) Push(n *Node) bool {
	if s.len == s.limiter {
		fmt.Printf("Stack is Full\n")
		return false
	}
	s.index++
	s.students[s.index] = n
	s.len++
	return true
}

func (s *Stack) Pop() (*Node, bool) {
	if s.len == 0 {
		fmt.Printf("Stack is Empty\n")
		return nil, false
	}
	n := s.students[s.index]
	s.index--
	s.len--
	return n, true
}

func (s *Stack) RenewStack(i int) bool {
	newlimit := s.limiter + i
	newstack := make([]*Node, newlimit)
	copy(newstack, s.students)
	s.students = newstack
	s.limiter = newlimit

	return true
}
