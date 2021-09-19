package singleton

import (
	"sync"
)

// 单例模式接口，导出的
type Singleton interface {
	foo()
}

// 单例模式实现类，包私有的
type singleton struct{}

func (s singleton) foo() {}

var (
	instance *singleton
	once     sync.Once
)

// 通过once.Do方法实现实例只被初始化一次
func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
