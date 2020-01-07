package main

type Counter struct {
	count int
	name  string
}

var instance *Counter

func GetInstance() *Counter {
	if instance == nil {
		instance = new(Counter)
	}
	return instance
}
func (s *Counter) AddOne() int {
	s.count++
	return s.count
}
