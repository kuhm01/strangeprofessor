package environment

type Minibuffer struct {
	LeftNode  *Node
	RightNode *Node
}

func NewMinibuffer() *Minibuffer {
	return &Minibuffer{}
}

func (mb *Minibuffer) Insert(node *Node) bool {
	if mb.LeftNode != nil {
		if mb.RightNode != nil {
			return false
		} else {
			mb.RightNode = &Node{Value: node.Value}
		}
	} else {
		mb.LeftNode = &Node{Value: node.Value}
	}

	return true
}

func (mb *Minibuffer) Operating(ch byte) (*Node, bool) {
	if mb.LeftNode == nil {
		return nil, false
	} else {
		if mb.RightNode == nil {
			return nil, false
		}
	}

	n1 := mb.LeftNode.Value
	mb.LeftNode = nil
	n2 := mb.RightNode.Value
	mb.RightNode = nil
	var rn int

	switch ch {
	case '+':
		rn = n1 + n2
	case '*':
		rn = n1 * n2
	default:
		return nil, false
	}

	return &Node{Value: rn}, true
}
