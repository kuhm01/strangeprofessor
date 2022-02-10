package environment

type StaticValues struct {
	Professor []*node
}

func NewProfessor() *StaticValues {
	return &StaticValues{Professor: make([]*node, 7)}
}

func (s *StaticValues) Request_Interview(index int, n *node) bool {
	i := index - 1
	s.Professor[i] = n
	return true
}

func (s *StaticValues) Response_Interview(index int) (*node, bool) {
	i := index - 1
	if s.Professor[i] != nil {
		return s.Professor[i], true
	} else {
		return nil, false
	}
}
