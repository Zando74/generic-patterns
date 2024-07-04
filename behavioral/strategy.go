package behavioral

type StrategyApplicant interface{}

type StrategyHandler[T StrategyApplicant] func(T)

type StrategyCreatorHandler[T StrategyApplicant] func() *T

type StrategyMethodHandler[T StrategyApplicant] func()

func NewStrategyHandler[T StrategyApplicant](handler func(T)) StrategyHandler[T] {
	return handler
}

type StrategyInstance[T StrategyApplicant] interface {
	Execute()
}
