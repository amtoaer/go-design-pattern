package prototype

type Prototype interface {
	clone() Prototype
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) clone() Prototype {
	stu := *s
	return &stu
}
