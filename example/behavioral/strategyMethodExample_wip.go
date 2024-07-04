package main

import "generic-patterns/behavioral"

type Navigator interface {
	behavioral.StrategyApplicant
}

type ConcreteNavigator struct {
	BuildRoute behavioral.StrategyMethodHandler[Navigator]
}

func (n *ConcreteNavigator) SetBuildRouteStrategy(strategy behavioral.StrategyMethodHandler[Navigator]) {
	n.BuildRoute = strategy
}

func (n *ConcreteNavigator) RoadStrategy() {
	println("Building route by road")
}

func (n *ConcreteNavigator) WalkStrategy() {
	println("Building route by walking")
}

func (n *ConcreteNavigator) PublicTransportStrategy() {
	println("Building route by public transport")
}

func MainStrategyMethodExample() {
	navigator := &ConcreteNavigator{}
	navigator.SetBuildRouteStrategy(navigator.RoadStrategy)

	navigator.BuildRoute()
}
