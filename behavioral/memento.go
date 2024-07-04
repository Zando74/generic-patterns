package behavioral

type Memento[T interface{}] struct {
	State T
}

func (m *Memento[T]) GetSavedState() T {
	return m.State
}

type Originator[T interface{}] struct {
	State T
}

func (e *Originator[T]) CreateMemento() *Memento[T] {
	return &Memento[T]{State: e.State}
}

func (e *Originator[T]) RestoreMemento(m *Memento[T]) {
	e.State = m.GetSavedState()
}

func (e *Originator[T]) SetState(state T) {
	e.State = state
}

func (e *Originator[T]) GetState() T {
	return e.State
}

type Caretaker[T interface{}] struct {
	MementoArray []*Memento[T]
}

func (c *Caretaker[T]) AddMemento(m *Memento[T]) {
	c.MementoArray = append(c.MementoArray, m)
}

func (c *Caretaker[T]) GetMemento(index int) *Memento[T] {
	return c.MementoArray[index]
}
