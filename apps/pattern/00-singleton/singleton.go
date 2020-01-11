package singleton

import "sync"

type Counter struct {
	count int
	name  string
	sync.RWMutex
}

var instance *Counter

func GetInstance() *Counter {
	if instance == nil {
		instance = new(Counter)
	}
	return instance
}
func (s *Counter) AddOne() int {
	s.Lock()
	defer s.Unlock()
	s.count++
	return s.count
}

func (s *Counter) getCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}
