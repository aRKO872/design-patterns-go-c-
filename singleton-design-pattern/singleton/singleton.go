package singleton

import (
	"log"
	"sync"
)

type Singleton struct {
}

var (
	singletonInstance *Singleton
	mutex             sync.Mutex
)

func NewSingleton() *Singleton {
	if singletonInstance == nil {
		mutex.Lock()
		if singletonInstance == nil {
			singletonInstance = &Singleton{}
			singletonInstance.Write(1)
		}
		mutex.Unlock()
	}
	return singletonInstance
}

func (s *Singleton) Write(ins int) {
	log.Printf("Singleton instance %d is being used\n", ins)
}
