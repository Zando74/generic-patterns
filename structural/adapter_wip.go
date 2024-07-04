package structural

type Adaptable interface{}
type Target interface{}

type Adapter[T Adaptable, U Target] struct {
	Adaptee *T
	Adapt   func(*T) *U
}

func (adapter *Adapter[T, U]) Adapted() *U {
	return adapter.Adapt(adapter.Adaptee)
}

func NewAdapter[T Adaptable, U Target](adaptee *T, handler func(*T) *U) *Adapter[T, U] {
	return &Adapter[T, U]{
		Adaptee: adaptee,
		Adapt:   handler,
	}
}
