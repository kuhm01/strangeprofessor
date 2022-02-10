package environment

type Environment struct {
	Student   *Stack
	Changed   *Queue
	Professor *StaticValues
}

func New() *Environment {
	return &Environment{
		Student:   NewStack(),
		Changed:   NewQueue(),
		Professor: NewProfessor(),
	}
}
