package behavioral

import "container/list"

type Observer interface {
	Notify(data interface{})
}

type Observable[T Observer] struct {
	Subs *list.List
}

func (o *Observable[T]) Subscribe(sub Observer) {
	o.Subs.PushBack(sub)
}

func (o *Observable[T]) Unsubscribe(sub Observer) {
	for s := o.Subs.Front(); s != nil; s = s.Next() {
		if s.Value == sub {
			o.Subs.Remove(s)
		}
	}
}

func (o *Observable[T]) Notify(data interface{}) {
	for s := o.Subs.Front(); s != nil; s = s.Next() {
		s.Value.(Observer).Notify(data)
	}
}

func NewObservable[T Observer]() Observable[T] {
	return Observable[T]{Subs: new(list.List)}
}
