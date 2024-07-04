package creational

import (
	"sync"
)

type Singleton[T any] struct {
	once            sync.Once
	instance        *T
	creationHandler func() *T
}

func NewSingleton[T any](creationHandler func() *T) *Singleton[T] {
	return &Singleton[T]{
		creationHandler: creationHandler,
	}
}

func (s *Singleton[T]) GetInstance() *T {
	s.once.Do(func() {
		instance := s.creationHandler()
		s.instance = instance
	})
	return s.instance
}
