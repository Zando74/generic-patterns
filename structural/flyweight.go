package structural

import "generic-patterns/behavioral"

type IntrinsicState interface{}

type Flyweight[T IntrinsicState] struct {
	Instances        map[string]*T
	CreationHandlers map[string]behavioral.StrategyCreatorHandler[T]
}

func (f *Flyweight[T]) GetInstance(key string) *T {

	existingInstance := f.Instances[key]

	if existingInstance == nil {
		// Should create it
		f.Instances[key] = f.CreationHandlers[key]()
	}

	return f.Instances[key]
}

func (f *Flyweight[T]) NewCreationHandler(key string, createFunc behavioral.StrategyCreatorHandler[T]) {
	f.CreationHandlers[key] = createFunc
}

func NewFlyweight[T IntrinsicState]() *Flyweight[T] {
	return &Flyweight[T]{
		Instances:        make(map[string]*T),
		CreationHandlers: make(map[string]behavioral.StrategyCreatorHandler[T]),
	}
}
