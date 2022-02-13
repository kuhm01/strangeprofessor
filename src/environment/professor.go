package environment

type StaticValues struct {
	Professor []*Node
}

func NewProfessor() *StaticValues {
	return &StaticValues{Professor: make([]*Node, 7)}
}

func (s *StaticValues) Setter(n int, index int) bool {
	i := index - 1
	node := s.Professor[i]
	if node == nil {
		node = &Node{}
	}
	node.Value += n
	s.Professor[i] = node
	return true
}

func (s *StaticValues) Request_Interview(index int, n *Node) bool {
	i := index - 1
	s.Professor[i] = n
	return true
}

func (s *StaticValues) Response_Interview(index int) (*Node, bool) {
	i := index - 1
	if s.Professor[i] != nil {
		return s.Professor[i], true
	} else {
		return nil, false
	}
}
