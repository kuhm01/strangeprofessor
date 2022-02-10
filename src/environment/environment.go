package environment

type Environment struct {
	Student *stack
}

func New() *Environment {
	return &Environment{Student: NewStack()}
}
