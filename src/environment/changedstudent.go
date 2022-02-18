package environment

import "fmt"

type Queue struct {
	students []*Node
	len      int
}

func NewQueue() *Queue {
	return &Queue{len: 0}
}

func (q *Queue) Set(n *Node) bool {
	node := &Node{Value: n.Value}
	q.students = append(q.students, node)
	q.len++
	return true
}

func (q *Queue) Get() (*Node, bool) {
	if q.len == 0 {
		fmt.Printf("Queue is Empty\n")
		return nil, false
	}
	item, items := q.students[0], q.students[1:]
	q.students = items
	q.len--
	return item, true
}
