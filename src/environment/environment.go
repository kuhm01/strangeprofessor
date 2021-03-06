package environment

import "fmt"

type Environment struct {
	Student    *Stack
	Changed    *Queue
	Professor  *StaticValues
	Buffer     *Buffer
	MiniBuffer *Minibuffer
}

func New() *Environment {
	return &Environment{
		Student:    NewStack(),
		Changed:    NewQueue(),
		Professor:  NewProfessor(),
		Buffer:     NewBuffer(),
		MiniBuffer: NewMinibuffer(),
	}
}

const (
	Int  = "int"
	Char = "char"
)

type PrintCharacter struct {
	Node      *Node
	PrintType string
}

type Buffer struct {
	Buffers    []*PrintCharacter
	FirstIndex int
	LastIndex  int
}

func NewBuffer() *Buffer {
	return &Buffer{Buffers: make([]*PrintCharacter, 20), FirstIndex: 0, LastIndex: 0}
}

func (b *Buffer) SetinBuffer(p *PrintCharacter) {
	if b.FirstIndex == b.LastIndex && b.Buffers[b.FirstIndex] != nil && b.Buffers[b.LastIndex] != nil {
		fmt.Printf("Buffer is Full\n")
	}
	tp := &PrintCharacter{Node: &Node{}}
	tp.Node.Value = p.Node.Value
	tp.PrintType = p.PrintType
	b.Buffers[b.LastIndex] = tp
	b.LastIndex++
	if b.LastIndex > 20 {
		b.LastIndex -= 20
	}
}

func (b *Buffer) GetinBuffer() *PrintCharacter {
	if b.FirstIndex == b.LastIndex {
		return nil
	}

	p := b.Buffers[b.FirstIndex]
	b.FirstIndex++
	if b.FirstIndex > 20 {
		b.FirstIndex -= 20
	}

	return &PrintCharacter{Node: p.Node, PrintType: p.PrintType}
}
