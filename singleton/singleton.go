package singleton

import "sync"

type Singleton struct{}

// 饿汉模式
var singleton = &Singleton{}

func Instance() *Singleton {
	return singleton
}

// 懒汉模式
var singletonLazy *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singletonLazy = &Singleton{}
	})
	return singletonLazy
}
