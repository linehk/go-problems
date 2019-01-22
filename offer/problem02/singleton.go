package problem02

import (
	"sync"
)

type example struct {
	name string
}

var instance *example
var once sync.Once

func GetInstance() *example {
	once.Do(func() {
		instance = new(example)
		instance.name = "第一次赋值单例"
	})
	return instance
}