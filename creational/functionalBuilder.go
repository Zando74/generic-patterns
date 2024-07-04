package creational

type Buildable interface{}

type FunctionalBuilderModifier[T Buildable] func(*T)

type FunctionalBuilder[T Buildable] struct {
	ToCraft T
	Actions []FunctionalBuilderModifier[T]
}

func (builder *FunctionalBuilder[T]) AddAction(action FunctionalBuilderModifier[T]) {
	builder.Actions = append(builder.Actions, action)
}

func (builder *FunctionalBuilder[T]) Build() T {
	toCraft := builder.ToCraft
	for _, action := range builder.Actions {
		action(&toCraft)
	}

	return toCraft
}

func (builder *FunctionalBuilder[T]) Reset() {
	builder.Actions = []FunctionalBuilderModifier[T]{}
}
