package creational

type FactoryApplicant interface{}

type Factory[T FactoryApplicant] struct {
	Make func() T
}

func NewFactory[T FactoryApplicant](make func() T) *Factory[T] {
	return &Factory[T]{Make: make}
}
