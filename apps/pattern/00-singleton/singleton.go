package singleton

import "sync"

type Counter struct {
	count int
	name  string
	sync.RWMutex
}

var instance *Counter
var mu sync.Mutex
var once sync.Once

func GetInstance() *Counter {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		instance = new(Counter)
	}
	return instance
}

func GetInstanceByOnce() *Counter {
	once.Do(func() {
		instance = new(Counter)
	})
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
