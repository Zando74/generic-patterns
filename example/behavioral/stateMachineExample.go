package main

import (
	"fmt"
	"generic-patterns/behavioral"
)

type Document struct {
	Title   string
	Content string
	// ...
	State *behavioral.StateMachine
}

func (d *Document) Moderate(approuved bool) error {
	if approuved {
		return d.State.GoTo(Approved)
	} else {
		return d.State.GoTo(Rejected)
	}
}

// Define the different possible states of the document
const (
	Draft behavioral.State = iota
	Moderation
	Approved
	Rejected
	Published
	MAX_BUILD_STATUS
)

// Define the mapping between state and their string representation
var DocumentStateToString = map[behavioral.State]string{
	Draft:      "Draft",
	Moderation: "Moderation",
	Approved:   "Approved",
	Rejected:   "Rejected",
	Published:  "Published",
}

// Define the transition rules between states
var DocumentStateTransitionRules = map[behavioral.State][]behavioral.State{
	Draft: {
		Moderation,
	},
	Moderation: {
		Approved, Rejected,
	},
	Rejected: {
		Draft,
	},
	Approved: {
		Published,
	},
}

func NewDocument(title, content string, state behavioral.State) (*Document, error) {

	documentStateMachine := (&behavioral.StateMachineBuilder{}).
		SetCurrentState(state).                            // Set the current state
		SetTransitionRules(&DocumentStateTransitionRules). // Set the transition rules between states
		SetStateToString(&DocumentStateToString).          // Set the mapping between state and string representation
		SetMaxUnreachableState(MAX_BUILD_STATUS).          // Set the maximum unreachable state (prevent invalid state value)
		Build()                                            // We can use dedicated builder Functionnal pattern to create a new StateMachine

	return &Document{
		Title:   title,
		Content: content,
		State:   &documentStateMachine,
	}, nil

}

func MainStateMachineExample() {

	draftDocument, _ := NewDocument("Draft Document", "This is a draft document", Draft) // You can start from any state

	draftDocument.State.GoTo(Moderation) // Transition from Draft to Moderation

	draftDocument.Moderate(true) // Transition from Moderation to Approved

	fmt.Println(draftDocument.State) // StateMachine implement Stringer Interface String() return currentState as string specified from mapping

	err := draftDocument.State.GoTo(Draft) // Trying to do invalid transition from Approved to Draft

	if err != nil {
		fmt.Println(err) // Output: Transition from Moderation to Published is not allowed
	}

}
