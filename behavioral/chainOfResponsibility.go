package behavioral

import "container/list"

type Query[D interface{}, R interface{}] struct {
	Data   D
	Result R
	Error  error
}

type Handler[D interface{}, R interface{}] interface {
	Handle(*Query[D, R])
}

type Broker[D interface{}, R interface{}] struct {
	Handlers list.List
}

func (b *Broker[D, R]) Subscribe(o Handler[D, R]) {
	b.Handlers.PushBack(o)
}

func (b *Broker[D, R]) Unsubscribe(o Handler[D, R]) {
	for s := b.Handlers.Front(); s != nil; s = s.Next() {
		if s.Value.(Handler[D, R]) == o {
			b.Handlers.Remove(s)
		}
	}
}

func (b *Broker[D, R]) Fire(q *Query[D, R]) {

	for s := b.Handlers.Front(); s != nil; s = s.Next() {
		s.Value.(Handler[D, R]).Handle(q)
		if q.Error != nil {
			return
		}
	}
}

func NewBroker[D interface{}, R interface{}]() *Broker[D, R] {
	return &Broker[D, R]{Handlers: list.List{}}
}
