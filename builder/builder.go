package builder

import "sync"

type Student struct {
	score map[string]int
	name  string
}

type builder struct {
	once    sync.Once
	student *Student
}

type BuilderInterface interface {
	WithScore(string, int) BuilderInterface
	WithName(string) BuilderInterface
	Build() *Student
}

var _ BuilderInterface = &builder{}

func Builder() BuilderInterface {
	return &builder{student: &Student{}}
}

func (b *builder) WithScore(subject string, score int) BuilderInterface {
	b.once.Do(func() {
		b.student.score = make(map[string]int)
	})
	b.student.score[subject] = score
	return b
}

func (b *builder) WithName(name string) BuilderInterface {
	b.student.name = name
	return b
}

func (b *builder) Build() *Student {
	return b.student
}
