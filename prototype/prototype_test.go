package prototype

import (
	"reflect"
	"testing"
)

func TestPrototype(t *testing.T) {
	first := &Student{
		Name: "amtoaer",
		Age:  20,
	}
	second := first.clone().(*Student)
	if !reflect.DeepEqual(first, second) {
		t.Error("原型模式出现错误")
	}
}
