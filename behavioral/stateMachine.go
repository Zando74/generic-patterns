package behavioral

import (
	"fmt"
	"strings"

	"github.com/Zando74/generic-patterns/creational"
)

type State int

func (state *State) String(stateToString *map[State]string) string {
	return (*stateToString)[*state]
}

type StateMachine struct {
	CurrentState        State
	TransitionRules     *map[State][]State
	stateToString       *map[State]string
	MaxUnreachableState State
}

type InvalidTransitionError struct {
	Current, Desired *State
	stateToString    *map[State]string
}

func (e *InvalidTransitionError) Error() string {

	return strings.ToUpper(fmt.Sprintf("Transition from %s to %s is Not Allowed", (*e.Current).String(e.stateToString), (*e.Desired).String(e.stateToString)))
}

type InvalidStateError struct{}

func (stateMachine *StateMachine) isTransitionAllowed(currentState, desiredState State) bool {
	for i := 0; i < len((*stateMachine.TransitionRules)[currentState]); i++ {
		if (*stateMachine.TransitionRules)[currentState][i] == desiredState {
			return true
		}
	}
	return false
}

func (stateMachine *StateMachine) GoTo(desiredState State) error {
	if stateMachine.isTransitionAllowed(stateMachine.CurrentState, desiredState) {
		stateMachine.CurrentState = desiredState
		return nil
	}
	return &InvalidTransitionError{&stateMachine.CurrentState, &desiredState, stateMachine.stateToString}
}

func (stateMachine *StateMachine) String() string {
	return stateMachine.CurrentState.String(stateMachine.stateToString)
}

type StateMachineBuilder struct {
	creational.FunctionalBuilder[StateMachine]
}

func (builder *StateMachineBuilder) SetCurrentState(initialState State) *StateMachineBuilder {

	builder.AddAction(func(stateMachine *StateMachine) {
		stateMachine.CurrentState = initialState
	})

	return builder
}

func (builder *StateMachineBuilder) SetTransitionRules(transitionRules *map[State][]State) *StateMachineBuilder {

	builder.AddAction(func(stateMachine *StateMachine) {
		stateMachine.TransitionRules = transitionRules
	})

	return builder
}

func (builder *StateMachineBuilder) SetStateToString(stateToString *map[State]string) *StateMachineBuilder {

	builder.AddAction(func(stateMachine *StateMachine) {
		stateMachine.stateToString = stateToString
	})

	return builder
}

func (builder *StateMachineBuilder) SetMaxUnreachableState(MaxUnreachableState State) *StateMachineBuilder {

	builder.AddAction(func(stateMachine *StateMachine) {
		stateMachine.MaxUnreachableState = MaxUnreachableState
	})

	return builder
}
