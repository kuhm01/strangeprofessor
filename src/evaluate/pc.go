package evaluate

import "fmt"

type Tag struct {
	IndexName string
	Index     int
}

type ProgramCounter struct {
	Index       int
	Tags        map[string]Tag
	PlusedIndex []int
	len         int
}

func NewPC(programLength int) *ProgramCounter {
	tgs := make(map[string]Tag)
	return &ProgramCounter{Index: 0, Tags: tgs, len: programLength}
}

func (pc *ProgramCounter) NewTag(name string) bool {
	t := &Tag{IndexName: name, Index: pc.Index}
	pc.Tags[name] = *t

	return true
}

func (pc *ProgramCounter) SetCounter(name string) bool {
	tag := pc.Tags[name]
	if tag.IndexName == name {
		pc.Index = tag.Index
		return true
	} else {
		fmt.Printf("Not equal flag name. got=%s", name)
		return false
	}
}

func (pc *ProgramCounter) Over() bool {
	if pc.Index < pc.len {
		return true
	} else {
		return false
	}
}

func (pc *ProgramCounter) isThisPlused() bool {
	for _, ii := range pc.PlusedIndex {
		if ii == pc.Index {
			return true
		}
	}

	return false
}

func (pc *ProgramCounter) ThisisPlused() {
	pc.PlusedIndex = append(pc.PlusedIndex, pc.Index)
}
