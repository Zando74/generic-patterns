package structural

import (
	"github.com/Zando74/generic-patterns/behavioral"
	"github.com/Zando74/generic-patterns/creational"
)

type NotInitializedExecutionHandler struct{}

func (e *NotInitializedExecutionHandler) Error() string {
	return "YOU MUST SET AN EXECUTION HANDLER BEFORE EXECUTING IT"
}

type Decorable interface{}

type Decorator[T Decorable] struct {
	*behavioral.Iterator[T]
	executionHandler behavioral.StrategyHandler[T]
}

func (d *Decorator[T]) Wrap(decorable T) *Decorator[T] {

	d.Extend(&behavioral.Iterator[T]{Item: &decorable})

	return d
}

func (d *Decorator[T]) SetExecutionHandler(handler *func(T)) *Decorator[T] {

	if d == nil {
		return nil
	}

	d.executionHandler = behavioral.NewStrategyHandler[T](*handler)
	return d
}

func (d *Decorator[T]) Execute() error {

	if d.executionHandler == nil {
		return &NotInitializedExecutionHandler{}
	}

	iterate := d.InitIterator()
	for it := iterate(); it != nil; it = iterate() {
		d.executionHandler(*it.Item)
	}

	return nil

}

func NewDecorator[T Decorable](decorable T) *Decorator[T] {
	return &Decorator[T]{
		Iterator: &behavioral.Iterator[T]{
			Item: &decorable,
		},
	}
}

func GenerateDecoratorFactory[T Decorable](handler *func(T), initialDecorable T, decorables ...T) *creational.Factory[*Decorator[T]] {

	factory := creational.NewFactory(func() *Decorator[T] {
		decorator := NewDecorator(initialDecorable)

		for _, decorable := range decorables {
			decorator = decorator.Wrap(decorable)
		}

		decorator = decorator.SetExecutionHandler(handler)
		return decorator
	})

	return factory
}
