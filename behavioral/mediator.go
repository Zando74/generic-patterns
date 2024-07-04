package behavioral

type Mediator interface{}

type Component[T Mediator] struct {
	Mediator *T
}

func (c *Component[T]) Register(m *T) {
	c.Mediator = m
}
