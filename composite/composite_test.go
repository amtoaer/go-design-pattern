package composite

import (
	"reflect"
	"testing"
)

func TestComposite(t *testing.T) {
	set := []struct {
		object Component
		result []string
	}{
		{&file{name: "file1"}, []string{"file:file1"}},
		{&folder{name: "folder1", components: []Component{&file{name: "file2"}}}, []string{"dir:folder1", "file:file2"}},
		{&folder{name: "folder2", components: []Component{&folder{name: "folder3", components: []Component{&file{name: "file3"}}}, &file{name: "file4"}}}, []string{"dir:folder2", "dir:folder3", "file:file3", "file:file4"}},
	}
	for _, testSet := range set {
		if !reflect.DeepEqual(testSet.object.Show(), testSet.result) {
			t.Error("组合模式出现错误")
		}
	}
}
