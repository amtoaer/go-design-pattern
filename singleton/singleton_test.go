package singleton

import "testing"

func TestSingleTon(t *testing.T) {
	t.Parallel()
	var (
		a = Instance()
		b = Instance()
	)
	if !(a == b && a != nil) {
		t.Error("饿汉模式出现错误")
	}
}

func TestSingleTonLazy(t *testing.T) {
	t.Parallel()
	var (
		a = GetInstance()
		b = GetInstance()
	)
	if !(a == b && a != nil) {
		t.Error("懒汉模式出现错误")
	}
}
