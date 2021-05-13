package composite

import "fmt"

type file struct {
	name string
}

type folder struct {
	components []Component
	name       string
}

type Component interface {
	Show() []string
}

func (f *file) Show() []string {
	return []string{fmt.Sprintf("file:%s", f.name)}
}

func (f *folder) Show() []string {
	var result []string
	result = append(result, fmt.Sprintf("dir:%s", f.name))
	for _, component := range f.components {
		result = append(result, component.Show()...)
	}
	return result
}
