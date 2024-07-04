package structural

type Composable interface{}

type Composite[T Composable] struct {
	Children []T
}

func (c *Composite[T]) Add(child T) {
	c.Children = append(c.Children, child)
}

func NewComposite[T Composable]() *Composite[T] {
	return &Composite[T]{}
}
