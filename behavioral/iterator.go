package behavioral

type Iterable interface{}

type Iterator[T Iterable] struct {
	Item *T
	next *Iterator[T]
}

func (i *Iterator[T]) HasNext() bool {
	return i.next != nil
}

func (i *Iterator[T]) SetNext(n *Iterator[T]) {
	i.next = n
}

func (i *Iterator[T]) InitIterator() func() *Iterator[T] {
	currentPtr := i

	return func() *Iterator[T] {
		defer func() {
			if currentPtr != nil {
				currentPtr = currentPtr.next
			}
		}()

		if currentPtr != nil {
			return currentPtr
		}
		return nil
	}
}

func (i *Iterator[T]) Last() *Iterator[T] {
	currentPtr := i
	for currentPtr.HasNext() {
		currentPtr = currentPtr.next
	}
	return currentPtr
}

func (i *Iterator[T]) Extend(n *Iterator[T]) {
	last := i.Last()
	last.SetNext(n)
}

func (i *Iterator[T]) Penultimate() *Iterator[T] {
	currentPtr := i
	for currentPtr.HasNext() && currentPtr.next.HasNext() {
		currentPtr = currentPtr.next
	}
	return currentPtr
}
