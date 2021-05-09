package singleton

import "sync"

type SingletonLazy struct{}

var singletonLazy *SingletonLazy
var once sync.Once

func GetInstance() *SingletonLazy {
	once.Do(func() {
		singletonLazy = &SingletonLazy{}
	})
	return singletonLazy
}
