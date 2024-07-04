package structural

type Proxifiable interface{}

type Proxy[T Proxifiable] struct {
	Wrapped          *T
	executionHandler func(*T)
}

func (p *Proxy[T]) Execute() error {
	p.executionHandler(p.Wrapped)
	return nil
}

func NewProxy[T Proxifiable](wrapped *T, handler func(*T)) *Proxy[T] {
	return &Proxy[T]{
		Wrapped:          wrapped,
		executionHandler: handler,
	}
}
