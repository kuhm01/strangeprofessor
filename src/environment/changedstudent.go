package environment

type Queue struct {
	students []*node
	len      int
}

func NewQueue() *Queue {
	return &Queue{len: 0}
}

func (q *Queue) Set(n *node) bool {
	q.students = append(q.students, n)
	return true
}

func (q *Queue) Get() (*node, bool) {
	if q.len == 0 {
		return nil, false
	}
	item, items := q.students[0], q.students[1:]
	q.students = items
	return item, true
}
