package environment

type Environment struct {
	Student *stack
	Changed *Queue
}

func New() *Environment {
	return &Environment{Student: NewStack(), Changed: NewQueue()}
}
