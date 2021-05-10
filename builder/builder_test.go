package builder

import (
	"reflect"
	"testing"
)

func TestBuilder(t *testing.T) {
	var (
		score = map[string]int{
			"语文": 100,
			"数学": 150,
			"英语": 120,
		}
		name = "李明"
	)
	builder := Builder().WithName(name)
	for k, v := range score {
		builder.WithScore(k, v)
	}
	a := builder.Build()
	b := &Student{
		name:  name,
		score: score,
	}
	if reflect.DeepEqual(a, b) == false {
		t.Error("建造者模式出现错误")
	}
}
