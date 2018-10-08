package M06_singleton

import "sync"

type single struct {
	O interface{}
}

var instantiated *single = nil
var once sync.Once

// New 非线程安全，单例模式
func New() *single {
	if instantiated == nil {
		instantiated = new(single)
	}
	return instantiated
}

// NewSafe 线程安全，单例模式
func NewSafe() *single {
	once.Do(func() {
		instantiated = &single{}
	})
	return instantiated
}


// =====================================

//
type Singleton struct {

}

var singleton *Singleton
var once2 sync.Once

func GetInstance() *Singleton{
	once2.Do(func(){
		singleton = &Singleton{}
	})
	return singleton
}