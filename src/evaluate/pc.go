package evaluate

type Tag struct {
	IndexName string
	Index     int
}

type ProgramCounter struct {
	Index int
	Tags  map[string]Tag
}

func NewPC() *ProgramCounter {
	tgs := make(map[string]Tag)
	return &ProgramCounter{Index: 0, Tags: tgs}
}

func (pc *ProgramCounter) NewTag(name string, index int) bool {
	t := &Tag{IndexName: name, Index: index}
	pc.Tags[name] = *t

	return true
}

func (pc *ProgramCounter) SetCounter(name string) bool {
	tag := pc.Tags[name]
	if tag.IndexName == name {
		pc.Index = tag.Index
		return true
	} else {
		return false
	}
}
