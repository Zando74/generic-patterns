package main

import (
	"fmt"
	"generic-patterns/behavioral"
)

type ConcreteState struct {
	Value string
}

func MainMementoExample() {

	caretaker := &behavioral.Caretaker[ConcreteState]{
		MementoArray: make([]*behavioral.Memento[ConcreteState], 0),
	}

	originator := &behavioral.Originator[ConcreteState]{
		State: ConcreteState{"A"},
	}

	fmt.Printf("Originator Current State: %s\n", originator.GetState().Value)
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState(ConcreteState{"B"})
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState(ConcreteState{"C"})
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.RestoreMemento(caretaker.GetMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.GetState())

	originator.RestoreMemento(caretaker.GetMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.GetState())

}
