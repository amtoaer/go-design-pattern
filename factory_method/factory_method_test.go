package factorymethod

import "testing"

func TestFactoryObject(t *testing.T) {
	t.Parallel()
	testCase := []Type{Start, End}
	factory := Factory{}
	for _, eventType := range testCase {
		if factory.Create(eventType).EventType() != eventType {
			t.Error("工厂类出现错误")
		}
	}
}

func TestFactoryMethod(t *testing.T) {
	t.Parallel()
	if OfStart().EventType() != Start || OfEnd().EventType() != End {
		t.Error("工厂方法出现错误")
	}
}
