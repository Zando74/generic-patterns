package structural

type Bridgeable interface{}

type Bridge[T Bridgeable] struct {
	Impl *T
}

func (b *Bridge[T]) SetImpl(impl T) {
	b.Impl = &impl
}
